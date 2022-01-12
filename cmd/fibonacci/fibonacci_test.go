package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	n   int
	res int
}

var testData = []TestItem{
	{5, 5},
	{6, 8},
}

func TestFibonacci(t *testing.T) {
	for _, v := range testData {
		t.Run(strconv.Itoa(v.n), func(t *testing.T) {
			res := fibonacciIter(v.n)
			assert.Equal(t, res, v.res)
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciIter(6)
	}
}
