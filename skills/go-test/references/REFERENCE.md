# Go Test Infrastructure Reference
This reference defines how the AI agent must use standardized test helpers when generating Go unit tests, integration tests, and infrastructure tests.

The goal is:
- Deterministic tests
- Reusable helpers
- Consistent structure
- High coverage
- Clean architecture alignment

## Optional but Recommended
Use the **Sequential Thinking MCP** when generating tests or infrastructure code. It helps the agent reason step-by-step before writing code, ensuring:
- Correct helper detection (`<root>/testutil`)
- Proper mock selection (SQL, Redis, HTTP, etc.)
- Clear success and failure coverage (≥ 80%)
- No duplication of template logic
- Deterministic and architecture-compliant tests
Not required for trivial pure functions, but highly recommended for non-trivial test generation.

## MANDATORY RULE: USE TEST HELPERS
Before generating any test code, the agent MUST first check whether the project already contains `<root>/testutil`.

### Existence Check (Required)
- If `<root>/testutil` exists:
	- Use existing helpers.
	- **Do NOT** regenerate templates.
	- **Do NOT** duplicate helper logic.

- If `<root>/testutil` does **NOT** exist:
	- Generate the test helper templates from [/templates](../templates/)
	- Create `<root>/testutil` using the canonical template files.
	- Then proceed with test generation.

This step is mandatory and must not be skipped.

Template test helper files include:
- `ent.go`
- `gin.go`
- `helper.go`
- `http.go`
- `redis.go`
- `sql.go`
These files are the canonical blueprint.

## SQL (database/sql + go-sqlmock)
If the code uses `database/sql`, the agent **MUST**:
- Use `testutil.NewSQLMock(t)`
- Not create `sqlmock.New()` inline
- Not manually call `ExpectationsWereMet()`
- Not manually close DB
- Use regexp matcher for SQL
- Escape PostgreSQL placeholders (`\$1`)

Example usage pattern:

```go
db, mock := testutil.NewSQLMock(t)
```

Transaction expectations:

```go
testutil.ExpectBegin(t, mock)
testutil.ExpectCommit(t, mock)
testutil.ExpectRollback(t, mock)
```

## Ent (ent + enttest)
If the code uses Ent ORM:
- Import `"<module>/ent/enttest"`
- Use `testutil.NewEntTestClient(t, dsn)`
- Prefer transaction rollback per test
- Do not manually run migrations
- Do not truncate tables manually

If transaction isolation helper exists:

```go
tx := testutil.WithTx(t, client)
repo := NewRepo(tx.Client())
```

## Redis (miniredis + go-redis)
If the code uses Redis:
- Use `testutil.NewRedisTestServer(t)`
- Do not connect to real Redis
- Do not use localhost
- Use `mr.FastForward()` for TTL testing
- Do not sleep to simulate expiration

Example:

```go
mr, client := testutil.NewRedisTestServer(t)
```

## Standardized Test Data Usage
All placeholder test data must be sourced from `<root>/testutil` constants.

Do not redefine literal strings inline. The agent must use predefined constants for:

- Names
- Emails
- Addresses
- Phone numbers
- Numeric values
- UUIDs (valid and invalid)
- URLs (valid and invalid)
- Fixed time
- Pagination

### Example Usage
Instead of:

```go
id := "550e8400-e29b-41d4-a716-446655440000"
email := "alice@fake.test"
web := "https://example.test"
phone := "+628111111111"
```

Use:

```go
id := testutil.ValidUUID1
email := testutil.EmailAlice
web := testutil.URLWeb
phone := testutil.PhoneAlice
```

## Context, Assertion, and Table Test Helpers
The agent must use standardized test utilities from `<root>/testutil` for context handling, assertions, and table-driven execution.

Do not reimplement these patterns inline.

### Context Helper
If the tested function requires `context.Context`, the agent **MUST** use:
```go
ctx := testutil.NewContext(t)
```

Do not use:

```go
context.Background()
context.TODO()
```

#### Example
Instead of:

```go
ctx := context.Background()
err := service.Process(ctx)
require.NoError(t, err)
```

Use:

```go
ctx := testutil.NewContext(t)
err := service.Process(ctx)
testutil.RequireNoError(t, err)
```

#### Why
- Ensures timeout safety
- Prevents hanging tests
- Automatically cleans up via `t.Cleanup`
- Improves determinism

### Assertion Helper Policy
All assertions in tests must use helpers from `testutil`.

Do not call `assert.*` or `require.*` directly inside test files. If an assertion helper does not exist in `testutil`, you may add a new helper following the same pattern (`t.Helper()` + wrapped testify call).

#### Example
```go
func TestCreateUser(t *testing.T) {
    ctx := testutil.NewContext(t)

    user, err := service.Create(ctx, testutil.NameAlice)

    testutil.RequireNoError(t, err)
    testutil.RequireNotNill(t, user)

    testutil.AssertEqual(t, testutil.NameAlice, user.Name)
    testutil.AssertNotNil(t, user.ID)
}
```

#### Strict Rule
The agent **MUST**:
- Use `testutil.Require*` for setup failures
- Use `testutil.Assert*` for value checks
- Never mix direct `assert.*` or `require.*` calls
- Add new helper functions to testutil if missing

#### When Adding a New Helper
If a needed assertion is not available:
- Add it to `testutil`
- Wrap `testify` call
- Add `t.Helper()`
- Follow naming convention

#### Why This Matters
Using assertion helpers ensures:
- Cleaner stack traces
- Centralized assertion behavior
- Consistent testing style
- Easier future refactor
- Enforcement across the codebase

### Table-Driven Test Helper
All table-driven tests should use:

```go
testutil.RunTableTests(t, tests)
```

Instead of manual loop:

```go
for _, tt := range tests {
    t.Run(tt.Name, func(t *testing.T) {
        tt.Run(t)
    })
}
```

#### Required Pattern
```go
func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "valid email with uppercase",
			input:       "  ALICE@FAKE.TEST ",
			expected:    testutil.EmailAlice,
			expectError: false,
		},
		{
			name:        "empty email",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format",
			input:       "not-an-email",
			expected:    "",
			expectError: true,
		},
	}

	table := make([]testutil.TableTestCase, 0, len(tests))
	for _, tt := range tests {
		tc := tt
		table = append(table, testutil.TableTestCase{
			Name: tc.name,
			Run: func(t *testing.T) {
				// Act
				result, err := NormalizeEmail(tc.input)

				// Assert
				if tc.expectError {
					testutil.RequireError(t, err)
					testutil.RequireEmpty(t, result)
					return
				}

				testutil.RequireNoError(t, err)
				testutil.RequireNotEmpty(t, result)
				testutil.AssertEqual(t, tc.expected, result)
			},
		})
	}

	testutil.RunTableTests(t, table)
}
```

#### Parallel Safety
`RunTableTests` runs subtests in parallel automatically.

The agent must ensure:
- No shared mutable state
- No global variable mutation
- No shared mock reuse across cases
- If shared setup is required, initialize inside each test case.

## HTTP & Gin Testing
This section defines how the agent must generate HTTP and Gin handler tests.

Tests must be deterministic, isolated, and use standardized helpers from `<root>/testutil`.

**Do not** reimplement HTTP setup logic inline.

### Generic HTTP Handlers (net/http)
If the handler uses `net/http`, the agent **MUST**:
- Use `testutil.NewJSONRequest`
- Use `testutil.PerformRequest`
- Use `testutil.DecodeJSON`
- Not manually instantiate `httptest.NewRecorder`
- Not manually encode `JSON` body
- Not manually decode response `JSON`

#### Required Pattern
```go
req := testutil.NewJSONRequest(t, http.MethodPost, "/users", payload)
rr := testutil.PerformRequest(handler, req)

testutil.RequireEqual(t, http.StatusCreated, rr.Code)

resp := testutil.DecodeJSON[ResponseType](t, rr)
```

##### Anti-Patterns
The agent must **NOT**:
- Start a real HTTP server
- Use real network ports
- Use `http.ListenAndServe`
- Manually marshal/unmarshal JSON
- Inline `httptest.NewRecorder()` repeatedly

### Gin Handlers
If the handler uses Gin, the agent **MUST**:
- Use `testutil.NewGinEngine()`
- Use `testutil.PerformGinRequest`
- Use `testutil.NewJSONRequest`
- Keep Gin in test mode
- Not call `gin.Default()` in tests
- Not bind to real ports

#### Required Pattern
```go
engine := testutil.NewGinEngine()
engine.POST("/users", handler)

req := testutil.NewJSONRequest(t, http.MethodPost, "/users", payload)
rr := testutil.PerformGinRequest(t, engine, req)

testutil.RequireEqual(t, http.StatusCreated, rr.Code)

resp := testutil.DecodeJSON[ResponseType](t, rr)
```

### Middleware Testing (Gin)
If middleware is involved:
- Register middleware on the test engine
- Test both authorized and unauthorized cases
- Do not bypass middleware unless explicitly testing handler logic only

Example:

```go
engine.Use(AuthMiddleware())
```