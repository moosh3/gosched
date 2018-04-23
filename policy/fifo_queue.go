// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package policy

import (
	"time"
)

const defaultQueueBuffer = 1024

// Queue represents a fifo queue data structure using an in memory
// heap as its data store. The struct also stores the queues priority level
// for cpu scheduling, and the ticker time for one cpu cycle (the higher the
// priority, the shorter the cpu cycle will be)
type Queue struct {
	job []interface{}
	buf []byte
}

// NewQueue returns a pointer to a fifo queue data structure with a set priorityLevel
func NewQueue() *Queue {
	return &Queue{buf: defaultQueueBuffer}
}

type PriorityQueue struct {
	q            Queue
	tickInterval time.Duration
	timeout      time.Duration
}

func NewPriorityQueue(tickInterval, timeout time.Duration) *PriorityQueue {
	return &PriorityQueue{
		tickInterval: tickInterval,
		timeout:      timeout,
	}
}
