package service

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	ctx := context.Background()

	require.NoError(t, client.Ping(ctx).Err())

	t.Cleanup(func() {
		require.NoError(t, client.Close())
		mr.Close()
	})

	return mr, client
}

func TestCacheService_Set(t *testing.T) {
	mr, client := NewRedisTestServer(t)
	s := &CacheService{rdb: client}

	tests := []struct {
		name    string
		key     string
		val     string
		ttl     time.Duration
		wantErr bool
	}{
		{
			name:    "success - set value with ttl",
			key:     "test-key",
			val:     "test-value",
			ttl:     10 * time.Second,
			wantErr: false,
		},
		{
			name:    "success - set value with no ttl",
			key:     "test-key-no-ttl",
			val:     "test-value-no-ttl",
			ttl:     0,
			wantErr: false,
		},
		{
			name:    "success - overwrite existing key",
			key:     "test-key",
			val:     "new-value",
			ttl:     5 * time.Second,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Act
			err := s.Set(ctx, tt.key, tt.val, tt.ttl)

			// Assert
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			// Verify in Redis
			got, err := mr.Get(tt.key)
			require.NoError(t, err)
			assert.Equal(t, tt.val, got)

			if tt.ttl > 0 {
				ttl := mr.TTL(tt.key)
				assert.True(t, ttl > 0 && ttl <= tt.ttl)
			}
		})
	}
}
