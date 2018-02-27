package PriorityQueue

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"time"
)

func TestQueue_Push(t *testing.T) {
	item := struct {}{}
	q := Build()
	q.Push(item, 2)
	assert.Equal(t, q.Len(), 1)
}

func TestQueue_Pull(t *testing.T) {
	item := struct {}{}
	q := Build()
	q.Push(item, 2)
	gotItem, _ := q.Pull()
	assert.Equal(t, gotItem, item)
}

func TestPriority(t *testing.T) {
	item1 := 111
	item2 := 222
	q := Build()
	q.Push(item2, 2)
	q.Push(item1, 1)
	gotItem, _ := q.Pull()
	assert.Equal(t, gotItem, item1)
}

func TestPrioritize(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	item1 := 111
	item2 := 222

	outCh, _ := Prioritize(ch1, ch2)

	ch2 <- item2
	time.Sleep(time.Millisecond)
	ch1 <- item1

	firstGot := <-outCh
	assert.Equal(t, firstGot, item2)
}