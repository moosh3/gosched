// Package fifo_queue provides a queue interface that can be used as a Job queue
// with the use of a Job Scheduler
package policy

import (
	"time"

	"github.com/aleccunningham/mlfq/internal/heap"
)


// Queue represents a fifo queue data structure using an in memory
// heap as its data store. The struct also stores the queues priority level
// for cpu scheduling, and the ticker time for one cpu cycle (the higher the
// priority, the shorter the cpu cycle will be)
func Queue struct {
	priorityLevel 	uint64
	tickerLength 	time.Time
	// fifo imported from container/heap
	heap Heap
}

// NewQueue returns a pointer to a fifo queue data structure with a set priorityLevel
func NewQueue(pl uint64) *Queue {
	return &Queue{ priorityLevel: pl }
}

// WithTicker returns a fifo queue with an associated priorityLevel and tickerLength
func NewQueueWithTicker(priorityLevel, tickerLength uint64) *Queue {
	queue := NewQueue(priorityLevel)
	heap.Init(&queue)
	queue.tickerLength = tickerLength // one cpu cycle
	return &queue
}

// getQueueWithPriority
func getQueueWithPriority(priority int) *Queue
