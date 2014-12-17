package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomCancel(done chan struct{}) {
	go func() {
		max := rand.Intn(1000) + 1000
		<-time.After(time.Duration(max) * time.Millisecond)
		close(done)
	}()
}

// START OMIT
func Gen(done chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				i = 10
			case out <- i:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	randomCancel(done)
	for i := range Gen(done) {
		fmt.Println(i)
	}
}

// END OMIT
