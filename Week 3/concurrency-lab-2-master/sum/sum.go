package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum uint64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&sum, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
