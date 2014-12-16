package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// START OMIT
func Generator() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			out <- i
		}

	}()
	return out
}

func LargeFileReader(fname string) <-chan []byte {
	out := make(chan []byte)
	go func() {
		f, _ := os.Open(fname)
		defer f.Close()
		data, _ := ioutil.ReadAll(f)
		out <- data
	}()
	return out
}

// END OMIT

func main() {
	numbers := Generator()
	for num := range numbers {
		fmt.Println(num)
	}
	fmt.Printf("%s\n", <-LargeFileReader("src/test.txt"))
}
