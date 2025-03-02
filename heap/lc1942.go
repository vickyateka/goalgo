package heap

import (
	"container/heap"
	"sort"
)

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func smallestChair(times [][]int, t_idx int) int {
	n := len(times)

	for i := 0; i < n; i++ {
		times[i] = append(times[i], i)
	}
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})

	taken := &MinHeap{}
	free := &MinHeapInt{}
	heap.Init(taken)
	heap.Init(free)

	currentSeat := 0

	for i := 0; i < n; i++ {
		at, lt, idx := times[i][0], times[i][1], times[i][2]

		for taken.Len() > 0 && (*taken)[0][0] <= at {
			heap.Push(free, (*taken)[0][1])
			heap.Pop(taken)
		}

		if free.Len() == 0 {
			if idx == t_idx {
				return currentSeat
			}
			heap.Push(taken, []int{lt, currentSeat})
			currentSeat++
		} else {
			x := heap.Pop(free).(int)
			if idx == t_idx {
				return x
			}
			heap.Push(taken, []int{lt, x})
		}
	}

	return -1
}

//func main(){
//	times := [][]int{{33889,98676},{80071,89737},{44118,52565},{52992,84310},{78492,88209},{21695,67063},{84622,95452},{98048,98856},{98411,99433},{55333,56548},{65375,88566},{55011,62821},{48548,48656},{87396,94825},{55273,81868},{75629,91467}}
//	t_idx := 6
//	smallestChair(times, t_idx)
//}
