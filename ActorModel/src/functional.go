package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func GenerateANumber() int {
	return r.Intn(10)
}

func PrintNum(a int) int {
	fmt.Println(a)
	return a
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

// START OMIT
func main() {
	num := GenerateANumber()

	PrintNum(num)

	num = Add(num, 5)
	num = Subtract(num, 3)

	PrintNum(num)

	//Or if you like Lisp
	PrintNum(Subtract(Add(PrintNum(GenerateANumber()), 5), 4))
}

// END OMIT
