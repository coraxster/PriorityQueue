package PriorityQueue

import (
	"container/heap"
	"sync"
)

type Queue struct {
	qh *QueueHeap
	m sync.Mutex
}

func Build() Queue {
	qh := make(QueueHeap, 0)
	heap.Init(&qh)
	queue := Queue{qh:&qh}
	return queue
}

func (q *Queue) Push (i interface{}, pr int) (bool, error){
	hi := HeapItem{
		priority: pr,
		data:i,
	}
	q.m.Lock()
	heap.Push(q.qh, &hi)
	q.m.Unlock()
	return true, nil
}

func (q *Queue) Pull() (interface{}, error) {
	q.m.Lock()
	item := heap.Pop(q.qh).(*HeapItem)
	q.m.Unlock()
	return item.data, nil
}

func (q *Queue) Len() int {
	return  q.qh.Len()
}


func Prioritize(ins... chan interface{}) (chan interface{}, error)  {
	out := make(chan interface{})
	q := Build()

	for pr, ch := range ins {
		go func(ch chan interface{}, pr int) {
			for item := range ch {
				q.Push(item, pr)
			}
		}(ch, pr)
	}
	go func() {
		for{
			if q.Len() > 0 {
				i, _ := q.Pull()
				out <- i
			}
		}
	}()
	return out, nil
}