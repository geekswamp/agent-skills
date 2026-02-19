package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

//
// ================================
// Redis Test Setup (miniredis)
// ================================
//

// NewRedisTestServer starts an in-memory Redis server and
// returns a ready-to-use Redis client.
func NewRedisTestServer(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	t.Helper()

	// Start in-memory Redis server
	mr, err := miniredis.Run()
	RequireNoError(t, err)

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	ctx := NewContext(t)

	RequireNoError(t, client.Ping(ctx).Err())

	t.Cleanup(func() {
		RequireNoError(t, client.Close())
		mr.Close()
	})

	return mr, client
}
