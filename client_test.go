package go_redislock

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var opt = &redis.Options{
	Addr:     "localhost:6379",
	Password: "a123456",
	DB:       0,
}

var lockKey = "test"

func TestLock_fail(t *testing.T) {
	client := Instance(opt)
	locked := client.Lock(lockKey, 10*time.Second)
	assert.True(t, locked)
	unLock := client.Lock(lockKey, 10*time.Second)
	assert.False(t, unLock)
	client.UnLock(lockKey)
}

func TestLock_success(t *testing.T) {
	client := Instance(opt)
	locked := client.Lock(lockKey, 10*time.Second)
	assert.True(t, locked)
	client.UnLock(lockKey)
	locked = client.Lock(lockKey, 10*time.Second)
	assert.True(t, locked)
	client.UnLock(lockKey)
}
