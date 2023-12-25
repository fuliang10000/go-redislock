# go-redislock

[![Go](https://img.shields.io/badge/Go->=1.21-green)](https://go.dev)

## 介绍
使用go编写的基于Redis实现的分布式锁，lockClient是单例模式，避免资源浪费，锁的操作是线程安全的

## 快速开始

### 安装
```bash
go get -u github.com/fuliang10000/go-redislock
```

### Use Demo
```go
package main

import (
	"context"
	"fmt"
	redisLock "github.com/fuliang10000/go-redislock"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	var opt = &redis.Options{
		Addr:     "localhost:6379",
		Password: "a123456",
		DB:       0,
	}
	// 获取一个单例客户端
	client := redisLock.Instance(context.Background(), redis.NewClient(opt))

	lockKey := "my_lock_key"
	// 获取锁
	locked := client.Lock(lockKey, 10*time.Second)
	if locked {
		// 释放锁
		defer client.UnLock(lockKey)
		// 执行业务逻辑
		fmt.Println("my work...")
	} else {
		panic("system is busy")
	}
}

```