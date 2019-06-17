package last_stone_weight

import (
	"container/heap"
)

type MaxHeap []int32

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int32))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Use Max heap to keep track the weightest stone
// of each move.
func lastStoneWeight(a []int32) int32 {
	m := MaxHeap(a)
	h := &m
	heap.Init(h)

	for h.Len() >= 2 {
		// Pick two most weightest stones,
		// calculate result and put the result back.
		max1 := heap.Pop(h).(int32)
		max2 := heap.Pop(h).(int32)
		result := max1 - max2
		heap.Push(h, result)
	}

	// The last root node in Max heap is the last stone weight.
	return heap.Pop(h).(int32)
}
