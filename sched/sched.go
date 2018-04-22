package mlfq

import (
	"time"

	"github.com/google/btree"
	"github.com/aleccunningham/mlfq/internal/os/proc"
)

// JobScheduler implements a CPU process scheduler, handing incoming Jobs and
// manipulating related priority queues accordingly.
type JobScheduler interface {
	Promote(Job, priority int) error
	PromoteAll([]Proc)
	Demote(Job, priority int)
	Recv(jobChan chan)
}

// mlfq is a JobScheduler implementation, scheduling Jobs on a virtualized CPU
type mlfq struct {
	mu     synx.Mutex
	q      []Queue // priority queues

	bt     *btree.BTree
}

// Promote takes a Job and moves it to a different queue that is of a higher priority than its current
func (m *mlfq) Promote(Job, priority int) error {}

// PromoteAll moves all Jobs to the highest priority queue
func (m *mlfq) PromoteAll([]Proc) {}

// Demote takes a Job and moves it to a different queue that is of a lower priority than its current
func (m *mlfq) Demote(Job, priority int) {}

// Recv takes a channel and will loop until all jobs in the channel are handled
func Recv(jobChan chan) {
	done := make(chan interface{})
  	defer close(done)

	for {
		select {
		case <-done:
			break loop
		case job := <-jobChan:
			queueJob(job)
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
