package util_test

import (
	"testing"

	"go-test-workspace/iteration-1/eval-6-context-data-mandatory/with_skill/outputs/testutil"
	"util"
)

func TestProcessUser(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		expectError bool
	}{
		{
			name:        "success with Alice email",
			email:       testutil.EmailAlice,
			expectError: false,
		},
		{
			name:        "success with Bob email",
			email:       testutil.EmailBob,
			expectError: false,
		},
		{
			name:        "success with John email",
			email:       testutil.EmailJohn,
			expectError: false,
		},
	}

	table := make([]testutil.TableTestCase, 0, len(tests))
	for _, tt := range tests {
		tc := tt
		table = append(table, testutil.TableTestCase{
			Name: tc.name,
			Run: func(t *testing.T) {
				// Arrange
				ctx := testutil.NewContext(t)

				// Act
				err := util.ProcessUser(ctx, tc.email)

				// Assert
				if tc.expectError {
					testutil.RequireError(t, err)
				} else {
					testutil.RequireNoError(t, err)
				}
			},
		})
	}

	testutil.RunTableTests(t, table)
}
