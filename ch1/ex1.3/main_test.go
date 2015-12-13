package main

import (
	"os"
	"strings"
	"testing"
)

// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func Echo1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}

func Echo2() {
	strings.Join(os.Args[1:], " ")
}
