package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"go-test-workspace/iteration-1/eval-2-repo-sql/with_skill/outputs/testutil"
)

func TestUserRepository_GetByID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		mockSetup   func(mock sqlmock.Sqlmock)
		expected    string
		expectError bool
	}{
		{
			name: "success",
			id:   testutil.ValidUUID1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(testutil.NameAlice)
				mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
					WithArgs(testutil.ValidUUID1).
					WillReturnRows(rows)
			},
			expected:    testutil.NameAlice,
			expectError: false,
		},
		{
			name: "user not found",
			id:   testutil.ValidUUID2,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
					WithArgs(testutil.ValidUUID2).
					WillReturnError(sql.ErrNoRows)
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "database error",
			id:   testutil.ValidUUID1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
					WithArgs(testutil.ValidUUID1).
					WillReturnError(errors.New("db error"))
			},
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
				// Arrange
				db, mock := testutil.NewSQLMock(t)
				repo := &UserRepository{db: db}
				tc.mockSetup(mock)

				// Act
				name, err := repo.GetByID(tc.id)

				// Assert
				if tc.expectError {
					testutil.AssertError(t, err)
					testutil.AssertEqual(t, tc.expected, name)
					return
				}

				testutil.AssertNoError(t, err)
				testutil.AssertEqual(t, tc.expected, name)
			},
		})
	}

	testutil.RunTableTests(t, table)
}
