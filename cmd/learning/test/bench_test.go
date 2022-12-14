package main

import "testing"

// cmd run
// go test -benchmem -run=^$ -count 5 -bench ^BenchmarkTest$ github.com/jamemyjamess/go-clean-architecture-demo/cmd/learning/test
func BenchmarkTest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Call your function
		TestErrorGroup()
	}
}
