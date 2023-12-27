package redislock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var opt = &redis.Options{
	Addr:     "localhost:6379",
	Password: "a123456",
	DB:       0,
}

var lockKey = "test"

// TestLock_success 期望获取锁成功
func TestLock_success(t *testing.T) {
	client := NewClient(context.Background(), redis.NewClient(opt))
	var wg sync.WaitGroup
	wg.Add(2)
	var locked bool
	go func() {
		defer wg.Done()
		defer client.UnLock(lockKey)
		locked = client.Lock(lockKey, 10*time.Second)
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		defer wg.Done()
		defer client.UnLock(lockKey)
		locked = client.Lock(lockKey, 10*time.Second)
	}()
	wg.Wait()
	assert.True(t, locked)
}

// TestLock_fail 期望获取锁失败
func TestLock_fail(t *testing.T) {
	client := NewClient(context.Background(), redis.NewClient(opt))
	var wg sync.WaitGroup
	wg.Add(2)
	var locked bool
	go func() {
		defer wg.Done()
		defer client.UnLock(lockKey)
		locked = client.Lock(lockKey, 10*time.Second)
		time.Sleep(50 * time.Millisecond)
	}()
	go func() {
		defer wg.Done()
		defer client.UnLock(lockKey)
		time.Sleep(10 * time.Millisecond)
		locked = client.Lock(lockKey, 10*time.Second)
	}()
	wg.Wait()
	assert.False(t, locked)
}
