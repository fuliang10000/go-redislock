package redislock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type Client struct {
	Ctx   context.Context
	redis *redis.Client
	mu    sync.Mutex
}

// NewClient 获取客户端
func NewClient(ctx context.Context, redis *redis.Client) *Client {
	return &Client{
		Ctx:   ctx,
		redis: redis,
	}
}

// Lock 获取锁
func (s *Client) Lock(lockKey string, expiration time.Duration) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	result, err := s.redis.SetNX(s.Ctx, lockKey, "locked", expiration).Result()
	if err != nil {
		panic(err)
	}
	return result
}

// UnLock 释放锁
func (s *Client) UnLock(lockKey string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.redis.Del(s.Ctx, lockKey)
}
