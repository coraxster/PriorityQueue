//Priority queue implementation, Heap based, Concurrency safe, FIFO
package PriorityQueue

import (
	"container/heap"
	"sync"
	//"reflect"
)

type Queue struct {
	qh *QueueHeap
	m sync.Mutex
	count uint64
}

func Build() Queue {
	qh := make(QueueHeap, 0)
	heap.Init(&qh)
	queue := Queue{qh:&qh}
	return queue
}

func (q *Queue) Push (i interface{}, pr int) (bool, error){
	q.m.Lock()
	defer q.m.Unlock()
	q.count++
	hi := HeapItem{
		order: q.count,
		priority: pr,
		data:i,
	}
	heap.Push(q.qh, &hi)
	return true, nil
}

func (q *Queue) Pull() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock();
	item := heap.Pop(q.qh).(*HeapItem)

	return item.data, nil
}

func (q *Queue) Len() int {
	q.m.Lock()
	defer q.m.Unlock()
	return  q.qh.Len()
}

//Receives channels in order of priority.(more is better) Returns output channel. (exp)
func Prioritize(ins... chan interface{}) (chan interface{}, error)  {
	out := make(chan interface{})
	q := Build()

	//go func() {
	//	var pushM sync.Mutex
	//	cases := make([]reflect.SelectCase, len(ins))
	//	for i, ch := range ins {
	//		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	//	}
	//	remaining := len(cases)
	//	for remaining > 0 {
	//		pushM.Lock()
	//		chosen, value, ok := reflect.Select(cases)
	//		if !ok {
	//			cases[chosen].Chan = reflect.ValueOf(nil)
	//			remaining -= 1
	//		} else {
	//			q.Push(value, chosen)
	//		}
	//		pushM.Unlock()
	//	}
	//
	//}()

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