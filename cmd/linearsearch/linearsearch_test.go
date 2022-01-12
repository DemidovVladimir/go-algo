package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestInput struct {
	sl  []int
	k   int
	res bool
}

var tdd = []TestInput{
	{[]int{1, 2, 3, 5}, 3, true},
	{[]int{1, 2, 3, 5}, 7, false},
}

func TestLinearSearch(t *testing.T) {
	for _, testCase := range tdd {
		t.Run(strconv.Itoa(testCase.k), func(t *testing.T) {
			res := linearSearch(testCase.sl, testCase.k)
			assert.Equal(t, res, testCase.res)
		})
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linearSearch([]int{1, 2, 3, 5}, 3)
	}
}
