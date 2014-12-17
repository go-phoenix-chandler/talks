package main

import (
	"fmt"
	"time"
)

func Start() <-chan int {
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

func End(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- (2 + i)
		}
	}()
	return out
}

// START OMIT
func Pipe(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- (1 + i)
		}
	}()
	return out
}

func main() {
	a := Start()
	b := Pipe(a)
	c := End(b)
	for i := range c {
		fmt.Println(i)
	}
}

// END OMIT
