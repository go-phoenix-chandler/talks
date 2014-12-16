package main

import (
	"fmt"
	"sync"
	"time"
)

func Start() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			out <- i
		}

	}()
	return out
}

func Worker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- (1 + i)
			time.Sleep(2 * time.Second)
		}
	}()
	return out
}

func createWorkers(nums <-chan int) []<-chan int {
	workers := make([]<-chan int, 0, 5)
	for i := 0; i < cap(workers); i++ {
		workers = append(workers, Worker(nums))
	}
	return workers
}

func closeChan(wg *sync.WaitGroup, out chan int) {
	wg.Wait()
	close(out)
}

// START OMIT
func FanIn(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go closeChan(&wg, out)
	return out
}

func main() {
	for o := range FanIn(createWorkers(Start())...) {
		fmt.Println(o)
	}
}

// END OMIT
