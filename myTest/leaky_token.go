package main

import (
	"fmt"
	"sync"
	"time"
)

type rateLimiter struct {
	lck      *sync.Mutex
	rate     float64   //最大速率限制
	balance  float64   //漏桶的剩余的空间
	capacity float64   //漏桶的最大容量限制
	lastTime time.Time //上次检查的时间
}

func NewRateLimiter(limitPerSecond, balance int) *rateLimiter {
	return &rateLimiter{
		lck:      new(sync.Mutex),
		rate:     float64(limitPerSecond),
		balance:  float64(balance),
		capacity: float64(balance),
		lastTime: time.Now(),
	}
}

func (r *rateLimiter) Check() bool {
	r.lck.Lock()
	defer r.lck.Unlock()

	ok := false
	now := time.Now()
	dur := now.Sub(r.lastTime).Seconds()
	water := dur * r.rate //计算这段时间内漏桶流出水的流量water
	r.lastTime = now
	r.balance += water //漏桶流出water容量的水，自然漏桶的余量多出water
	if r.balance > r.capacity {
		r.balance = r.capacity
	}
	if r.balance >= 1 { //漏桶余量足够容下当前的请求
		r.balance -= 1
		ok = true
	}
	return ok
}

//漏桶算法
func main() {
	limiter := NewRateLimiter(10, 1)
	start := time.Now()
	count := 0
	for i := 0; i < 1e3; i++ {
		if limiter.Check() {
			fmt.Println(i)
			count++
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Println("count:", count)
	fmt.Println(time.Now().Sub(start).Seconds())
}
