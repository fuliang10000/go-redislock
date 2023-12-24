package go_redislock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	instance *Singleton
	once     sync.Once
)

type Singleton struct {
	Ctx    context.Context
	Client *redis.Client
	mu     sync.Mutex
}

// Instance 获取单例客户端
func Instance(ctx context.Context, client *redis.Client) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Ctx:    ctx,
			Client: client,
		}
	})
	return instance
}

// Lock 获取锁
func (s *Singleton) Lock(lockKey string, expiration time.Duration) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	result, err := s.Client.SetNX(s.Ctx, lockKey, "locked", expiration).Result()
	if err != nil {
		panic(err)
	}
	return result
}

// UnLock 释放锁
func (s *Singleton) UnLock(lockKey string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Client.Del(s.Ctx, lockKey)
}
