package queue

import (
	"container/list"
	"fmt"
)

// Queue implemented using container/list with generics:
type Queue[T any] struct {
	data *list.List
}

// NewQueue creates and returns a new Queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: list.New()}
}

// Enqueue adds a value to the back of the queue
func (q *Queue[T]) Enqueue(value T) {
	q.data.PushBack(value)
}

// Dequeue removes and returns the value from the front of the queue
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("queue is empty")
	}
	front := q.data.Front()
	q.data.Remove(front)
	return front.Value.(T), nil
}

// Front returns the value at the front without removing it
func (q *Queue[T]) Front() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("queue is empty")
	}
	return q.data.Front().Value.(T), nil
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.data.Len() == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return q.data.Len()
}
