package gosched

import (
	"os"
	"time"
)

type Job struct {
	// jid is the position in a given slice of Jobs
	jid int
	// proc is the os process that the job is scheduled to run
	proc *os.Process

	// priority is the priority level that corresponds to the priority queues
	// in order to place the job on the right queue
	priority JobPriority
	// timeSlice can be used as a mock process by representing io/cpu
	// blocks of which the scheduler must handle
	timeSlice time.Time
	// state is the state of the current Job, set by the scheduler
	state JobStateType
}

// JobPriority separates the four job queues from one another, giving each queue a certain "weight" or
// priority on a scale of 0-4, with 0 being of the highest priority, and 4 as the last.
type JobPriority uint32

// MinPriority
func (p JobPriority) MinPriority() int { return 0 }

// MaxPriority
func (p JobPriority) MaxPriority() int { return 0 }

// JobStateType
type JobStateType int

const (
	JobStateNone    JobStateType = iota // unititialized
	JobStateReady                       // job is ready to be scheduled
	JobStateWait                        // job request is waiting
	JobStateRunning                     // job request is running its process
	JobStateDone                        // job request has finished
)

// MoveToReady
func MoveToReady(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Ready"
	job.process = proc
}

// MoveToWait
func MoveToWait(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Wait"
	job.process = proc
}

// MoveToRunning
func MoveToRunning(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Running"
	job.process = proc
}

// MoveToDone
func MoveToDone(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Done"
	job.process = proc
}

func findByPid(pid int) *os.Process { return nil }
