package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	h   []int
	res []int
}

var testData = []TestItem{
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{2, 5, 3, 7, 4}, []int{2, 3, 4, 5, 7}},
}

func TestBubbleSort(t *testing.T) {
	for _, v := range testData {
		t.Run("Bubble sort", func(t *testing.T) {
			res := bubbleSort(v.h)
			assert.Equal(t, res, v.res)
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort(rand.Perm(10000))
	}
}
