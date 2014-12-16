package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			out <- i
			time.Sleep(200 * time.Millisecond)
		}

	}()
	return out
}

// START OMIT
func Consumer(in <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range in {
			fmt.Println(i)
		}
	}()
}

func main() {
	nums := Producer()
	wg := &sync.WaitGroup{}
	Consumer(nums, wg)

	// Do Other work
	wg.Wait()
}

// END OMIT
