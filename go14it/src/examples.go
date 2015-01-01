package main

import "fmt"

type Languages int

//go:generate stringer -type=Languages
const (
	Irish Languages = iota
	Gaelic
	Welsh
	Gaeilge = Irish
)

func NewForRangeExample() {
	stuff := []string{"Happy", "New", "Year"}
	for range stuff {
		fmt.Println("doing stuff...")
	}
}
