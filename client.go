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
	*redis.Client
}

// Instance 获取redis客户端
func Instance(opt *redis.Options) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			redis.NewClient(opt),
		}
	})
	return instance
}

// Lock 获取锁
func (s *Singleton) Lock(lockKey string, expiration time.Duration) bool {
	ctx := context.Background()
	result, err := s.Client.SetNX(ctx, lockKey, "locked", expiration).Result()
	if err != nil {
		panic(err)
	}
	return result
}

// UnLock 释放锁
func (s *Singleton) UnLock(lockKey string) {
	ctx := context.Background()
	s.Client.Del(ctx, lockKey)
}
