package testutil

import (
	"testing"

	"entgo.io/ent/dialect"
	"<module>/ent"
	"<module>/ent/enttest"
	"<module>/ent/migrate"
)

// ================================
// Ent Test Client (Project Scoped)
// ================================

// NewEntTestClient creates a new ent test client.
// It runs migration automatically and registers cleanup.
func NewEntTestClient(t *testing.T, dsn string) *ent.Client {
	t.Helper()

	client := enttest.Open(
		t,
		dialect.Postgres,
		dsn,
		enttest.WithMigrateOptions(
			migrate.WithGlobalUniqueID(true),
		),
	)

	RequireNotNill(t, client)

	t.Cleanup(func() {
		RequireNoError(t, client.Close())
	})

	return client
}
