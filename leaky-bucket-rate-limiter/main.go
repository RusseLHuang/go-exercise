package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Limiter default by 5 seconds
type Limiter struct {
	StartTime  time.Time
	LimitTime  time.Duration
	MaxRequest uint64
	Count      uint64
}

func New(
	maxRequest uint64,
) Limiter {
	return Limiter{
		MaxRequest: maxRequest,
		LimitTime:  time.Duration(5 * time.Second),
	}
}

func (limiter *Limiter) Take() bool {
	countIncr := atomic.LoadUint64(&limiter.Count) + 1

	if limiter.Count == 0 || time.Now().After(limiter.StartTime.Add(limiter.LimitTime)) {
		limiter.StartTime = time.Now()
		limiter.Count = 0
	} else if countIncr > limiter.MaxRequest {
		return false
	}

	atomic.AddUint64(&limiter.Count, 1)

	return true
}

func main() {
	limiter := New(5)

	for i := 0; i < 6; i++ {
		fmt.Println(limiter.Take())
	}

	time.Sleep(time.Duration(5 * time.Second))

	for i := 0; i < 6; i++ {
		fmt.Println(limiter.Take())
	}

}
