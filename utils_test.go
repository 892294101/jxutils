package jxutils

import (
	"sync"
	"testing"
)

func BenchmarkToHashConcurrent(t *testing.B) {
	concurrency := 1000000 // 并发数
	wg := sync.WaitGroup{}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = ToHash(string("sfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdfsfksajdfkjasdf"))
		}()
	}
	wg.Wait()
}
