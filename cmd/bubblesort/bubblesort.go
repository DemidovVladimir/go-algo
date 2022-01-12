package main

func bubbleSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			if n[j] > n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}

	return n
}
