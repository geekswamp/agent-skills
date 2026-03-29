package util

import (
	"testing"

	"go-test-workspace/iteration-1/eval-6-context-data-mandatory/without_skill/testutil"
)

func TestProcessUser(t *testing.T) {
	tests := []struct {
		name  string
		email string
	}{
		{
			name:  "process alice email",
			email: testutil.EmailAlice,
		},
		{
			name:  "process bob email",
			email: testutil.EmailBob,
		},
		{
			name:  "process john email",
			email: testutil.EmailJohn,
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
				err := ProcessUser(ctx, tc.email)

				// Assert
				testutil.AssertNoError(t, err)
			},
		})
	}

	testutil.RunTableTests(t, table)
}
