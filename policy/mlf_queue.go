// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package policy

import (
	"log"
	"sync"

	"github.com/google/btree"
	"github.com/sirupsen/logrus"
)

type MlfqScheduler struct {
	sync.Mutex
	q []Queue // priority queues

	bt *btree.BTree
}

func NewMlfqScheduler() *MlfqScheduler {
	return &MlfqScheduler{}
}

// Promote takes a Job and moves it to a different queue that is of a higher priority than its current
func (m *mlfq) Promote(Job, priority int) error {}

// PromoteAll moves all Jobs to the highest priority queue
func (m *mlfq) PromoteAll([]Proc) {}

// Demote takes a Job and moves it to a different queue that is of a lower priority than its current
func (m *mlfq) Demote(Job, priority int) {}

// Recv takes a channel and will loop until all jobs in the channel are handled
func Recv(jobChan chan struct{}) {
	done := make(chan interface{})
	defer close(done)

	for {
		select {
		case <-done:
			break
		case job := <-jobChan:
			queueJob(job)
		}
	}
}

const (
	DefaultQueueCount = 4
)

func init() {
	for _, queue := range DefaultQueueCount {
		q := NewJobQueue()
		if err := q.Subscribe(); err != nil {
			logrus.Fatal(err)
		}
	}
}

// queueJob places a given job onto a retrieved queue based upon the
// priority level of the job. The job will now sit on the queue for one CPU cycle.
func (m mlfq) queueJob(j Job) error {
	priority, error := jobPriority(Job)
	if err != nil {
		log.Fatalf("Failed to get job")
	}
	queue := getQueueWithPriority(priority)
	queue.add(j)
}
