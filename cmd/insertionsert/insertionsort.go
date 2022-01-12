package main

func insertionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := i; j > 0 && n[j-1] > n[j]; j-- {
			n[j-1], n[j] = n[j], n[j-1]
		}
	}

	return n
}
