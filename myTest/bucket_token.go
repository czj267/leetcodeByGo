package main

//令牌桶算法
import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const TOKEN_GRANULARITY = 1000

type MAP struct {
	lock   sync.RWMutex
	bucket map[string]*tokenBucket
}

//是否可实现一种无锁算法 TODO
func (m *MAP) Set(k string, times, interval int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.bucket[k]; !ok {
		tb := new(tokenBucket)
		tb.Init(times, interval)
		m.bucket[k] = tb
	}
	return
}
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func NewMAP() *MAP {
	return &MAP{bucket: make(map[string]*tokenBucket)}
}

type tokenBucket struct {
	lastQuery time.Time
	tokens    int
	capacity  int
	step      int
	add       int
	mu        sync.Mutex
}

func (tb *tokenBucket) Init(quota int, interval int) {
	tb.capacity = quota * TOKEN_GRANULARITY                  //最多保留的令牌
	tb.tokens = quota * TOKEN_GRANULARITY                    //当前保留的令牌
	tb.step = quota * TOKEN_GRANULARITY / quota              //每次需消耗的令牌
	tb.add = quota * TOKEN_GRANULARITY * interval / interval //每秒增加的令牌
	tb.lastQuery = time.Now()
}

func (tb *tokenBucket) TokenBucketQuery() error {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	dur := now.Sub(tb.lastQuery)
	token := int(dur.Nanoseconds()/1000000000) * tb.add
	tb.lastQuery = now
	tb.tokens += token           //增加此段时间的令牌
	if tb.tokens > tb.capacity { //超过最大令牌数重置
		tb.tokens = tb.capacity
	}
	if tb.tokens >= tb.step { // 与每次消耗的做对比
		tb.tokens -= tb.step
		return nil
	}

	return errors.New("Not enough")
}

//@test
func main() {
	var tb tokenBucket
	tb.Init(5, 1) //1s内调用5次
	cnt := 0
	for {
		err := tb.TokenBucketQuery()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("take")
		}
		cnt += 1
		fmt.Println(cnt)
		time.Sleep(100000 * time.Microsecond)
		//time.Sleep(5* time.Second)
	}
}
