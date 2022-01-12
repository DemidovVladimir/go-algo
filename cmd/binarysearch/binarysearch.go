package main

func binarySearch(h []int, k int) bool {
	low := 0
	high := len(h) - 1

	for low <= high {
		median := (low + high) / 2

		if h[median] < k {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(h) || h[low] != k {
		return false
	}

	return true
}

func main() {}
