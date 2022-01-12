package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	type TestItem struct {
		h   []int
		res []int
	}

	l := 1000
	r := rand.Perm(l)
	s := make([]int, l)
	copy(s, r)
	sort.Ints(s)

	testData := []TestItem{
		{r, s},
		{[]int{2, 5, 3, 7, 4, 8, 1}, []int{1, 2, 3, 4, 5, 7, 8}},
	}

	for _, v := range testData {
		t.Run("Quick sort", func(t *testing.T) {
			res := quickSort(v.h, 0, len(v.h)-1)
			assert.Equal(t, res, v.res)
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quickSort(rand.Perm(1000), 0, 999)
	}
}
