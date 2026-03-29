package service

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestCacheService_Set(t *testing.T) {
	db, mock := redismock.NewClientMock()
	service := &CacheService{
		rdb: db,
	}

	ctx := context.Background()
	key := "test_key"
	val := "test_value"
	ttl := 10 * time.Minute

	t.Run("success", func(t *testing.T) {
		mock.ExpectSet(key, val, ttl).SetVal("OK")

		err := service.Set(ctx, key, val, ttl)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("redis error", func(t *testing.T) {
		mock.ExpectSet(key, val, ttl).SetErr(assert.AnError)

		err := service.Set(ctx, key, val, ttl)
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
