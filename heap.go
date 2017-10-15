package PriorityQueue

import "container/heap"

type HeapItem struct {
	order        int
	index    int
	priority int
	data     interface{}
}

// A QueueHeap implements heap.Interface and holds Items.
type QueueHeap []*HeapItem

func (pq QueueHeap) Len() int {
	return len(pq)
}

func (pq QueueHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority == pq[j].priority {
		return pq[i].order < pq[j].order
	}
	return pq[i].priority < pq[j].priority
}

func (pq QueueHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *QueueHeap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*HeapItem)
	item.index = n
	item.order = n
	*pq = append(*pq, item)
}

func (pq *QueueHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *QueueHeap) update(item *HeapItem, value string, priority int) {
	item.data = value
	item.priority = priority
	heap.Fix(pq, item.index)
}