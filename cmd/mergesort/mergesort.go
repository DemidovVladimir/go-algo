package main

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	first := mergeSort(a[:len(a)/2])
	second := mergeSort(a[len(a)/2:])

	return merge(first, second)
}

func merge(first, second []int) []int {
	final := make([]int, len(first)+len(second))

	i := 0
	j := 0
	r := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final[r] = first[i]
			r++
			i++
		} else {
			final[r] = second[j]
			r++
			j++
		}
	}

	for i < len(first) {
		final[r] = first[i]
		r++
		i++
	}

	for j < len(second) {
		final[r] = second[j]
		r++
		j++
	}

	return final
}

func main() {}
