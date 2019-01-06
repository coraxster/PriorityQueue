//Priority queue implementation, Heap based, Concurrency safe, FIFO
package PriorityQueue

import (
	"container/heap"
	"sync"
	//"reflect"
	"errors"
	"math"
)

type Queue struct {
	qh    *QueueHeap
	m     sync.Mutex
	count uint64
}

func Build() Queue {
	qh := make(QueueHeap, 0)
	heap.Init(&qh)
	queue := Queue{qh: &qh}
	return queue
}

func (q *Queue) Push(i interface{}, pr int) (bool, error) {
	q.m.Lock()
	defer q.m.Unlock()
	if uint64(q.qh.Len()) == math.MaxUint64 {
		return false, errors.New("queue is full")
	}
	if q.count == math.MaxUint64 {
		q.count = q.qh.CollapseOrder()
	}
	q.count++
	hi := HeapItem{
		order:    q.count,
		priority: pr,
		data:     i,
	}
	heap.Push(q.qh, &hi)
	return true, nil
}

func (q *Queue) Pull() (interface{}, error) {
	if q.Len() == 0 {
		return nil, errors.New("empty")
	}
	q.m.Lock()
	defer q.m.Unlock()
	item := heap.Pop(q.qh).(*HeapItem)
	return item.data, nil
}

func (q *Queue) Len() int {
	q.m.Lock()
	defer q.m.Unlock()
	return q.qh.Len()
}

//Receives channels in order of priority.(more is better) Returns output channel. (exp)
func Prioritize(ins ...chan interface{}) (chan interface{}, error) {
	out := make(chan interface{})
	q := Build()

	c := sync.NewCond(&q.m)

	for pr, ch := range ins {
		go func(ch chan interface{}, pr int) {
			for item := range ch {
				q.Push(item, pr)
				c.Signal()
			}
		}(ch, pr)
	}
	go func() {
		for {
			q.m.Lock()
			if q.qh.Len() == 0 {
				c.Wait()
			}
			item := heap.Pop(q.qh).(*HeapItem).data
			q.m.Unlock()

			out <- item
		}
	}()
	return out, nil
}
