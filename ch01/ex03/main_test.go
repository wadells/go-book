package main

import (
	"strconv"
	"testing"
)

const SIZE = 100

func makeStringRange(n int) []string {
	a := make([]string, n)
	for i := range a {
		a[i] = strconv.Itoa(i)
	}
	return a
}

func BenchmarkJoins(b *testing.B) {
	args := makeStringRange(SIZE)
	for i := 0; i < b.N; i++ {
		printStringJoin(args)
	}
}

func BenchmarkConcat(b *testing.B) {
	args := makeStringRange(SIZE)
	for i := 0; i < b.N; i++ {
		printStringConcat(args)
	}
}
