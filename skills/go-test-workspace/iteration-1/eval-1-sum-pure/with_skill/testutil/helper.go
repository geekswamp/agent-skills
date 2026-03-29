package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ================================
// Context Helpers
// ================================

// NewContext returns a background context with timeout for tests.
func NewContext(t *testing.T) context.Context {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	t.Cleanup(cancel)

	return ctx
}

// ================================
// Assertion Helpers
// ================================

// AssertContains is a shortcut for assert.Contains with helper mark.
func AssertContains(t *testing.T, s any, contains any) bool {
	t.Helper()
	return assert.Contains(t, s, contains)
}

// AssertEqual is a shortcut for assert.Equal with helper mark.
func AssertEqual(t *testing.T, expected any, actual any) bool {
	t.Helper()
	return assert.Equal(t, expected, actual)
}

// AssertSame is a shortcut for assert.Same with helper mark.
func AssertSame(t *testing.T, expected any, actual any) bool {
	t.Helper()
	return assert.Same(t, expected, actual)
}

// AssertEmpty is a shortcut for assert.Empty with helper mark.
func AssertEmpty(t *testing.T, object any) bool {
	t.Helper()
	return assert.Empty(t, object)
}

// AssertNil is a shortcut for assert.Nil with helper mark.
func AssertNil(t *testing.T, object any) bool {
	t.Helper()
	return assert.Nil(t, object)
}

// AssertTrue is a shortcut for assert.True with helper mark.
func AssertTrue(t *testing.T, value bool) bool {
	t.Helper()
	return assert.True(t, value)
}

// AssertFalse is a shortcut for assert.False with helper mark.
func AssertFalse(t *testing.T, value bool) bool {
	t.Helper()
	return assert.False(t, value)
}

// AssertError is a shortcut for assert.Error with helper mark.
func AssertError(t *testing.T, err error) bool {
	t.Helper()
	return assert.Error(t, err)
}

// AssertNoError is a shortcut for assert.NoError with helper mark.
func AssertNoError(t *testing.T, err error) bool {
	t.Helper()
	return assert.NoError(t, err)
}

// AssertNotNil is a shortcut for assert.NotNil with helper mark.
func AssertNotNil(t *testing.T, object any) bool {
	t.Helper()
	return assert.NotNil(t, object)
}

// AssertGreater is a shortcut for assert.Greater with helper mark.
func AssertGreater(t *testing.T, a any, b any) bool {
    t.Helper()
    return assert.Greater(t, a, b)
}


// RequireError is a shortcut for require.Error with helper mark.
func RequireError(t *testing.T, err error) {
	t.Helper()
	require.Error(t, err)
}

// RequireNoError is a shortcut for require.NoError with helper mark.
func RequireNoError(t *testing.T, err error) {
	t.Helper()
	require.NoError(t, err)
}

// RequireNotNil is a shortcut for require.NotNil with helper mark.
func RequireNotNil(t *testing.T, object any) {
	t.Helper()
	require.NotNil(t, object)
}

// RequireSame is a shortcut for require.Same with helper mark.
func RequireSame(t *testing.T, expected any, actual any) {
	t.Helper()
	require.Same(t, expected, actual)
}

// RequirePanicsWithError is a shortcut for require.PanicsWithError with helper mark.
func RequirePanicsWithError(t *testing.T, errString string, f assert.PanicTestFunc) {
	t.Helper()
	require.PanicsWithError(t, errString, f)
}

// RequireLen is a shortcut for require.Len with helper mark.
func RequireLen(t *testing.T, object any, length int) {
	t.Helper()
	require.Len(t, object, length)
}

// RequireNotEmpty is a shortcut for require.NotEmpty with helper mark.
func RequireNotEmpty(t *testing.T, object any) {
	t.Helper()
	require.NotEmpty(t, object)
}

// RequireEqual is a shortcut for require.Equal with helper mark.
func RequireEqual(t *testing.T, expected any, actual any) {
	t.Helper()
	require.Equal(t, expected, actual)
}


// RequireFalse is a shortcut for require.False with helper mark.
func RequireFalse(t *testing.T, value bool) {
	t.Helper()
	require.False(t, value)
}

// RequireTrue is a shortcut for require.True with helper mark.
func RequireTrue(t *testing.T, value bool) {
	t.Helper()
	require.True(t, value)
}

// RequireEmpty is a shortcut for require.Empty with helper mark.
func RequireEmpty(t *testing.T, object any) {
	t.Helper()
	require.Empty(t, object)
}

// RequireContains is a shortcut for require.Contains with helper mark.
func RequireContains(t *testing.T, s any, contains any) {
	t.Helper()
	require.Contains(t, s, contains)
}

// RequireGreater is a shortcut for assert.Greater with helper mark.
func RequireGreater(t *testing.T, a any, b any) {
    t.Helper()
    require.Greater(t, a, b)
}

// ================================
// Table Test Helper
// ================================

// TableTestCase represents a table-driven test case.
type TableTestCase struct {
	Name string
	Run  func(t *testing.T)
}

// RunTableTests runs table-driven tests safely.
func RunTableTests(t *testing.T, tests []TableTestCase) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			tt.Run(t)
		})
	}
}

// ================================
// Test Data (Deterministic)
// ================================

const (
	// ===============================
	// Names
	// ===============================

	NameAlice = "Alice"
	NameBob   = "Bob"
	NameJohn  = "John Doe"
	NameJean  = "Jean Doe"
	NameJanet = "Janet"

	// ===============================
	// Emails
	// ===============================
	
	EmailAlice = "alice@fake.test"
	EmailBob   = "bob@fake.test"
	EmailJohn  = "john.doe@fake.test"
	EmailJean  = "jean.doe@test.fake"
	EmailJanet = "janet@fake.test"

	// ===============================
	// Addresses
	// ===============================

	AddressAlice = "123 Main Street, Suite 101, Test City, TC 12345"
	AddressBob   = "456 Example Avenue, Apt 22B, Sample Town, ST 67890"
	AddressJohn  = "789 Mockingbird Lane, Demo District, DD 10101"
	AddressJean  = "321 Placeholder Road, Unit 9, Fakeville, FV 20202"
	AddressJanet = "654 Integration Blvd, Floor 3, Dev City, DC 30303"

	// ===============================
	// Phone Numbers
	// ===============================

	PhoneAlice = "+628111111111"
	PhoneBob   = "+628122222222"
	PhoneJohn  = "+15551234567"
	PhoneJean  = "+441234567890"
	PhoneJanet = "+819012345678"

	// ===============================
	// Numeric Test Values
	// ===============================

	NumberZero       = 0
	NumberOne        = 1
	NumberTen        = 10
	NumberHundred    = 100
	NumberThousand   = 1000
	NegativeOne      = -1
	LargeNumber      = 999999
	MaxLimit         = 100
	MinLimit         = 1

	// ===============================
	// UUIDs
	// ===============================

	ValidUUID1   = "550e8400-e29b-41d4-a716-446655440000"
	ValidUUID2   = "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	InvalidUUID1 = "not-a-uuid"
	InvalidUUID2 = "550e8400e29b41d4a716446655440000" // missing hyphens

	// ===============================
	// URLs (Valid)
	// ===============================

	URLWeb      = "https://example.test"
	URLAPI      = "https://api.example.test/v1/users"
	URLCDN      = "https://cdn.example.test/assets/image.png"
	URLAppLogin = "https://app.example.test/login"

	// ===============================
	// URLs (Invalid)
	// ===============================

	InvalidURL1 = "htp://invalid-url"     // wrong scheme
	InvalidURL2 = "www.example.test"      // missing scheme
	InvalidURL3 = "https:/broken-url.com" // malformed scheme
	InvalidURL4 = "://missing-scheme.com" // missing protocol
	InvalidURL5 = "not-a-url"

	// ===============================
	// Timeout
	// ===============================

	Timeout = 3 * time.Second

	// ===============================
	// Pagination
	// ===============================
	
	PageOne  = 1
	PageTwo  = 2
	PageSize = 10
	MaxPageSize = 100
)

// ================================
// Time Helpers
// ================================

// FixedTime returns deterministic time for testing.
func FixedTime() time.Time {
	return time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
}
