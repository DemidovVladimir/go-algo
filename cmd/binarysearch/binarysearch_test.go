package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	h   []int
	k   int
	res bool
}

var testData = []TestItem{
	{[]int{1, 2, 3, 4, 5}, 3, true},
	{[]int{1, 2, 3, 4, 5}, 7, false},
}

func TestBinarySearch(t *testing.T) {
	for _, v := range testData {
		t.Run(strconv.Itoa(v.k), func(t *testing.T) {
			res := binarySearch(v.h, v.k)
			assert.Equal(t, res, v.res)
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarySearch([]int{1, 2, 3, 4, 5}, 3)
	}
}
