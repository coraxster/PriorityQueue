# PriorityQueue [![Build Status](https://travis-ci.org/coraxster/PriorityQueue.svg?branch=master)](https://travis-ci.org/coraxster/PriorityQueue)

Priority Queue implementation on Go.

1. On top of native heap
2. Multithread access safe (using mutex)
3. Fifo


## Usage
```go
//make queue
q := PriorityQueue.Build()

//push
pr: = 1
q.Push(something.(interface{}), pr)

//pull
somethingWithHightPr := q.Pull()
```

## Usage, prioritize channels 
```go
//ins - []chan interface{}
out := PriorityQueue.Prioritize(ins...)
```

[example](https://github.com/coraxster/PriorityQueueExp)
