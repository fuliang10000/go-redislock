package go_redislock

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

func TestLock_success(t *testing.T) {
	client := Instance(context.Background(), redis.NewClient(opt))
	var wg sync.WaitGroup
	wg.Add(2)
	var locked bool
	go func() {
		defer wg.Done()
		locked = client.Lock(lockKey, 10*time.Second)
		client.UnLock(lockKey)
	}()
	go func() {
		defer wg.Done()
		locked = client.Lock(lockKey, 10*time.Second)
		client.UnLock(lockKey)
	}()
	wg.Wait()
	assert.True(t, locked)
}

func TestLock_fail(t *testing.T) {
	client := Instance(context.Background(), redis.NewClient(opt))
	var wg sync.WaitGroup
	wg.Add(2)
	var locked bool
	go func() {
		defer wg.Done()
		locked = client.Lock(lockKey, 10*time.Second)
		time.Sleep(50 * time.Millisecond)
		client.UnLock(lockKey)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		locked = client.Lock(lockKey, 10*time.Second)
		client.UnLock(lockKey)
	}()
	wg.Wait()
	assert.False(t, locked)
}
