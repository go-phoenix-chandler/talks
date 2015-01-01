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

// NewForRangeExample prints a string for each element in the stuff array
func NewForRangeExample() {
	stuff := []string{"Happy", "New", "Year"}
	for range stuff {
		fmt.Println("doing stuff...")
	}
}
