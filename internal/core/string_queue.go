package core

import (
	"container/list"
	"fmt"
)

type StringQueue struct {
	list *list.List
}

func NewStringQueue(values []string) *StringQueue {
	q := &StringQueue{list.New()}
	for _, value := range values {
		q.list.PushBack(value)
	}
	return q
}

func (q *StringQueue) Dequeue() string {
	front := q.list.Front()
	value := front.Value
	q.list.Remove(front)
	return fmt.Sprint(value)
}

func (q *StringQueue) DequeueOrDefault(def string) string {

	if q.list.Len() == 0 {
		return def
	}

	front := q.list.Front()
	value := front.Value
	q.list.Remove(front)
	return fmt.Sprint(value)
}

func (q *StringQueue) Enqueue(value string) {
	q.list.PushBack(value)
}

func (q *StringQueue) Empty() bool {
	return q.list.Len() == 0
}
