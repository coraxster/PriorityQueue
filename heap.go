package PriorityQueue

type HeapItem struct {
	order        uint64
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
}

func (pq *QueueHeap) Push(x interface{}) {
	item := x.(*HeapItem)
	*pq = append(*pq, item)
}

func (pq *QueueHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq QueueHeap) CollapseOrder() uint64 {
	orderMap := make(map[uint64][]*HeapItem)
	orders := make([]uint64, 0)
	for _, hi := range pq {
		orderMap[hi.order] = append(orderMap[hi.order], hi)
		orders = append(orders, hi.order)
	}

	for i, order := range orders {
		for _, hi := range orderMap[order] {
			hi.order = uint64(i+1)
		}
	}
	return uint64(len(orders))
}
