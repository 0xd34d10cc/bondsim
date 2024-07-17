package sim

import (
	"container/heap"
	"time"
)

type taskFn = func()

type Task struct {
	fn taskFn
	at time.Time
}

type queue []Task

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].at.Before(q[j].at)
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x any) {
	item := x.(Task)
	*q = append(*q, item)
}

func (q *queue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = Task{} // avoid memory leak
	*q = old[0 : n-1]
	return item
}

type Scheduler interface {
	Schedule(t time.Time, fn taskFn)
}

type scheduler struct {
	q queue
}

func newScheduler() *scheduler {
	return &scheduler{
		q: make(queue, 0),
	}
}

func (s *scheduler) Schedule(t time.Time, fn taskFn) {
	heap.Push(&s.q, Task{at: t, fn: fn})
}

func (s *scheduler) RunUntil(to time.Time) {
	for {
		if len(s.q) == 0 {
			break
		}

		task := heap.Pop(&s.q).(Task)
		if task.at.After(to) {
			heap.Push(&s.q, task)
			break
		}

		task.fn()
	}
}
