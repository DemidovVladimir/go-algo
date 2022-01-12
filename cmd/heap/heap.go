package main

type MaxHeap struct {
	sl []int
}

func (h *MaxHeap) Insert(key int) {
	h.sl = append(h.sl, key)
	h.maxHeapifyUp(len(h.sl) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {

}

func main() {

}
