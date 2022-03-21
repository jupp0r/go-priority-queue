// Package pq implements a priority queue data structure on top of container/heap.
// As an addition to regular operations, it allows an update of an items priority,
// allowing the queue to be used in graph search algorithms like Dijkstra's algorithm.
// Computational complexities of operations are mainly determined by container/heap.
// In addition, a map of items is maintained, allowing O(1) lookup needed for priority updates,
// which themselves are O(log n).
package pq

import (
	"container/heap"
	"errors"
)

// PriorityQueue represents the queue
type PriorityQueue[T any] struct {
	itemHeap *itemHeap[T]
	lookup   map[interface{}]*item[T]
}

// New initializes an empty priority queue.
func New[T any]() PriorityQueue[T] {
	return PriorityQueue[T]{
		itemHeap: &itemHeap[T]{},
		lookup:   make(map[interface{}]*item[T]),
	}
}

// Len returns the number of elements in the queue.
func (p *PriorityQueue[T]) Len() int {
	return p.itemHeap.Len()
}

// Insert inserts a new element into the queue. No action is performed on duplicate elements.
func (p *PriorityQueue[T]) Insert(v T, priority float64) {
	_, ok := p.lookup[v]
	if ok {
		return
	}

	newItem := &item[T]{
		value:    v,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
	p.lookup[v] = newItem
}

// Pop removes the element with the highest priority from the queue and returns it.
// In case of an empty queue, an error is returned.
func (p *PriorityQueue[T]) Pop() (T, error) {
	if len(*p.itemHeap) == 0 {
		var zeroVal T
		return zeroVal, errors.New("empty queue")
	}

	item := heap.Pop(p.itemHeap).(*item[T])
	delete(p.lookup, item.value)
	return item.value, nil
}

// UpdatePriority changes the priority of a given item.
// If the specified item is not present in the queue, no action is performed.
func (p *PriorityQueue[T]) UpdatePriority(x T, newPriority float64) {
	item, ok := p.lookup[x]
	if !ok {
		return
	}

	item.priority = newPriority
	heap.Fix(p.itemHeap, item.index)
}

type itemHeap[T any] []*item[T]

type item[T any] struct {
	value    T
	priority float64
	index    int
}

func (ih *itemHeap[T]) Len() int {
	return len(*ih)
}

func (ih *itemHeap[T]) Less(i, j int) bool {
	return (*ih)[i].priority < (*ih)[j].priority
}

func (ih *itemHeap[T]) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap[T]) Push(x interface{}) {
	it := x.(*item[T])
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap[T]) Pop() interface{} {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
