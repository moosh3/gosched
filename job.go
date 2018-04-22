package gosched

import (
	"os"
	"time"
)

type Job struct {
	jid     int         // jid is the position in a given slice of Jobs
	pid     int         // pid of the Jobs associated process
	process *os.Process // proc is the os process that the job is scheduled to run

	priority  JobPriority // corresponds to the priority queues in order to place the job on the right queue
	timeSlice time.Time   // time io/mem blocks of which the scheduler must handle
	state     JobStateType
}

// JobPriority separates the four job queues from one another, giving each queue a certain "weight" or
// priority on a scale of 0-4, with 0 being of the highest priority, and 4 as the last.
type JobPriority uint32

// MinPriority returns the min priority provided by the scheduler, which is the same
// value as the priority on the lowest priority FIFO queue
func (p JobPriority) MinPriority() int { return 0 }

// MaxPriority returns the max priority provided by the scheduler, which is the same
// value as the priority on the highest priority FIFO queue
func (p JobPriority) MaxPriority() int { return 0 }

// JobStateType is an integer representation of the current state a specific Job
type JobStateType int

const (
	JobStateNone    JobStateType = iota // unititialized
	JobStateReady                       // job is ready to be scheduled
	JobStateWait                        // job request is waiting
	JobStateRunning                     // job request is running its process
	JobStateDone                        // job request has finished
)

// MoveToReady changes a Jobs state to Ready
func MoveToReady(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Ready"
	job.process = proc
}

// MoveToWait changes a Jobs state to Ready
func MoveToWait(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Wait"
	job.process = proc
}

// MoveToRunning changes a Jobs state to Ready
func MoveToRunning(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Running"
	job.process = proc
}

// MoveToDone changes a Jobs state to Ready
func MoveToDone(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Done"
	job.process = proc
}

func findByJid(jid int) *os.Process { return nil }

func findByPid(pid int) *os.Process { return nil }
