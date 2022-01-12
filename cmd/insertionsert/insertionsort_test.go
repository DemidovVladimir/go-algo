package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	type TestItem struct {
		h   []int
		res []int
	}

	r := rand.Perm(10000)
	s := sort.IntSlice(r)

	var testData = []TestItem{
		{r, s},
		{[]int{2, 5, 3, 7, 4}, []int{2, 3, 4, 5, 7}},
	}

	for _, v := range testData {
		t.Run("Insertion sort", func(t *testing.T) {
			res := insertionSort(v.h)
			assert.Equal(t, res, v.res)
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort(rand.Perm(10000))
	}
}
