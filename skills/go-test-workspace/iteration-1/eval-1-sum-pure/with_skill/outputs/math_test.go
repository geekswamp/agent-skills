package math

import (
	"testing"

	"go-test-workspace/iteration-1/eval-1-sum-pure/with_skill/testutil"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        testutil.NumberTen,
			b:        testutil.NumberTen,
			expected: 20,
		},
		{
			name:     "negative numbers",
			a:        testutil.NegativeOne,
			b:        -5,
			expected: -6,
		},
		{
			name:     "mixed numbers",
			a:        testutil.NumberTen,
			b:        testutil.NegativeOne,
			expected: 9,
		},
		{
			name:     "zero values",
			a:        testutil.NumberZero,
			b:        testutil.NumberZero,
			expected: 0,
		},
		{
			name:     "large numbers",
			a:        testutil.LargeNumber,
			b:        testutil.NumberOne,
			expected: 1000000,
		},
	}

	table := make([]testutil.TableTestCase, 0, len(tests))
	for _, tt := range tests {
		tc := tt
		table = append(table, testutil.TableTestCase{
			Name: tc.name,
			Run: func(t *testing.T) {
				// Act
				result := Sum(tc.a, tc.b)

				// Assert
				testutil.AssertEqual(t, tc.expected, result)
			},
		})
	}

	testutil.RunTableTests(t, table)
}
