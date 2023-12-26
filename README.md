# go-redislock

[![Go](https://img.shields.io/badge/Go->=1.21-green)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/fuliang10000/go-redislock.svg)](https://github.com/fuliang10000/go-redislock/releases)
[![Report](https://goreportcard.com/badge/github.com/fuliang10000/go-redislock)](https://goreportcard.com/report/github.com/fuliang10000/go-redislock)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/fuliang10000/go-redislock)
[![License](https://img.shields.io/github/license/fuliang10000/go-redislock)](https://github.com/fuliang10000/go-redislock/blob/main/LICENSE)

## 介绍
> 使用go编写的基于Redis实现的分布式锁

## 快速开始

### 安装
```bash
go get -u github.com/fuliang10000/go-redislock
```

### 测试
```bash
go test github.com/fuliang10000/go-redislock  -v -cover
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
	// 获取客户端
	client := redisLock.NewClient(context.Background(), redis.NewClient(opt))

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