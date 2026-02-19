package testutil

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

//
// ================================
// SQLMock Setup
// ================================
//

// NewSQLMock creates a new sql.DB and sqlmock instance.
// It automatically registers cleanup.
func NewSQLMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
		sqlmock.MonitorPingsOption(true),
	)
	
	RequireNoError(t, err)

	t.Cleanup(func() {
		RequireNoError(t, mock.ExpectationsWereMet())
		RequireNoError(t, db.Close())
	})

	return db, mock
}


//
// ================================
// SQL Result Helpers
// ================================
//

// NewResult creates a mocked sql.Result.
func NewResult(lastInsertID int64, rowsAffected int64) sql.Result {
	return sqlmock.NewResult(lastInsertID, rowsAffected)
}

//
// ================================
// Transactions Helpers
// ================================
//

// ExpectBegin registers an expected BEGIN transaction call.
func ExpectBegin(t *testing.T, mock sqlmock.Sqlmock) {
	t.Helper()
	mock.ExpectBegin()
}

// ExpectCommit registers an expected COMMIT transaction call.
func ExpectCommit(t *testing.T, mock sqlmock.Sqlmock) {
	t.Helper()
	mock.ExpectCommit()
}

// ExpectRollback registers an expected ROLLBACK transaction call.
func ExpectRollback(t *testing.T, mock sqlmock.Sqlmock) {
	t.Helper()
	mock.ExpectRollback()
}

//
// ================================
// Connection Helpers
// ================================
//

// ExpectPing expects a database ping call.
func ExpectPing(t *testing.T, mock sqlmock.Sqlmock) {
	t.Helper()
	mock.ExpectPing()
}