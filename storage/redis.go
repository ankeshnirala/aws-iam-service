package storage

import (
	"context"
	"os"

	"github.com/ankeshnirala/go/aws-iam-service/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap/zapcore"
)

type RedisStore struct {
	db *redis.Client
}

func NewRedisStore() (*RedisStore, error) {
	options := &redis.Options{Addr: os.Getenv("REDISDB_URL")}

	rdb := redis.NewClient(options)

	// Ping the Redis server to check the connectivity.
	_, err := rdb.Ping(context.TODO()).Result()

	if err != nil {
		logger.Fatal(err.Error(), zapcore.Field{Type: zapcore.SkipType})
	}

	logger.Info("Connected to RedisDB!", zapcore.Field{Type: zapcore.SkipType})

	defer rdb.Close()
	return &RedisStore{db: rdb}, nil
}
