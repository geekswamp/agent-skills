package testutil

import (
	"testing"

	"<your-module>/ent/enttest"
	"<your-module>/ent/migrate"
)

//
// ================================
// Ent Test Client (Project Scoped)
// ================================
//

// NewEntTestClient creates a new ent test client.
// It runs migration automatically and registers cleanup.
func NewEntTestClient(t *testing.T, dsn string) *enttest.Client {
	t.Helper()

	client := enttest.Open(
		t,
		"postgres",
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
