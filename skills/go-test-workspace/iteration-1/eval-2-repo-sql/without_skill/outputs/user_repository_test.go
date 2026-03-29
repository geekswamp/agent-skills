package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &UserRepository{db: db}

	t.Run("success", func(t *testing.T) {
		id := "123"
		expectedName := "John Doe"

		rows := sqlmock.NewRows([]string{"name"}).AddRow(expectedName)
		mock.ExpectQuery("SELECT name FROM users WHERE id = \\?").
			WithArgs(id).
			WillReturnRows(rows)

		name, err := r.GetByID(id)
		if err != nil {
			t.Errorf("error was not expected: %s", err)
		}
		if name != expectedName {
			t.Errorf("expected %s, got %s", expectedName, name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		id := "456"

		mock.ExpectQuery("SELECT name FROM users WHERE id = \\?").
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		_, err := r.GetByID(id)
		if err != sql.ErrNoRows {
			t.Errorf("expected sql.ErrNoRows, got %s", err)
		}
	})
}
