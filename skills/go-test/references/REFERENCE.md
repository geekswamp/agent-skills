# Go Test Infrastructure Reference
This reference defines how the AI agent must use standardized test helpers when generating Go unit tests, integration tests, and infrastructure tests.

The goal is:
- Deterministic tests
- Reusable helpers
- Consistent structure
- High coverage
- Clean architecture alignment

## MANDATORY RULE: USE TEST HELPERS
When generating tests, the agent **MUST**:
- Prefer existing test helper utilities over inline setup.
- Never duplicate test setup logic that already exists in `<root>/testutil`.
- Always use deterministic placeholder data defined in `testutil`.
- Avoid creating ad-hoc mocks when helper templates are available.
- If a reusable template exists, use it instead of re-implementing setup logic.
- All infrastructure templates are located in [templates](../templates/). Test helper files: `ent.go`, `helper.go`, `redis.go`, and `sql.go`.

The agent must:
- Check the [templates](../templates/) directory before generating new helper patterns.
- Use the existing template structure as the canonical implementation.
- Not recreate template logic inside test files.
- Not redefine helpers already provided in `<root>/testutil`.

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
- UUIDs (valid and invalid)
- URLs (valid and invalid)
- Fixed time

### Example Usage
Instead of:

```go
id := "550e8400-e29b-41d4-a716-446655440000"
email := "alice@fake.test"
web := "https://example.test"
```

Use:

```go
id := testutil.ValidUUID1
email := testutil.EmailAlice
web := testutil.URLWeb
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

### Assertion Helper
For fatal setup assertions, use:

```go
testutil.RequireNoError(t, err)
```

Instead of:

```go
require.NoError(t, err)
```

#### When to Use
- DB setup
- Mock creation
- Transaction begin
- Repository initialization
- Critical preconditions

Do not replace value assertions (`assert.Equal`) unless helper exists.

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
func TestDivide(t *testing.T) {
	tests := []struct {
		Name string
		Run  func(t *testing.T)
	}{
		{
			Name: "success",
			Run: func(t *testing.T) {
				result, err := Divide(10, 2)
				testutil.RequireNoError(t, err)
				require.Equal(t, 5, result)
			},
		},
		{
			Name: "division by zero",
			Run: func(t *testing.T) {
				_, err := Divide(10, 0)
				require.Error(t, err)
			},
		},
	}

	testutil.RunTableTests(t, tests)
}
```

#### Parallel Safety
`RunTableTests` runs subtests in parallel automatically.

The agent must ensure:
- No shared mutable state
- No global variable mutation
- No shared mock reuse across cases
- If shared setup is required, initialize inside each test case.