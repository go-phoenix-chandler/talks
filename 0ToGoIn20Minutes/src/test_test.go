// test is an example package containing the most basic
// tests possible.
package presentation

import (
	"testing"
)

// Double doubles the provided value
func Double(n int) int {
	return n * 2
}

// Fibonacci will what you'd think
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// TestDouble verifies that Double returns 2 * n
func TestDouble(t *testing.T) {
	if Double(2) != 4 {
		t.Error("Double is not doubling...")
	}

	if Double(3) == 5 {
		t.Error("ERROR: 3 doubled doens't equal 5")
	}
}

// BenchmarkFibonacci will provide benchmark statistics for
// the Fibonacci function
func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(10)
	}
}
