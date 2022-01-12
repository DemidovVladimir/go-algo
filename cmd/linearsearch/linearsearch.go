package main

func linearSearch(sl []int, k int) bool {
	for _, v := range sl {
		if v == k {
			return true
		}
	}
	return false
}

func main() {
}
