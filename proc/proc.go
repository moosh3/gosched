package proc

import (
	"os"
	"time"
)

// Proc is a wrapper around os.Process, which is a struct containing information in regards to os.StartProcess,
// defined in the os.ProcAttr struct (and in the Proc struct). Proc defines the timeSlice, allowing the Scheduler
// to allocate the proper time slot for the process to run. The Process is run with a context instance for
// passing on information regarding the ProcessState. A separate JobState is also defined to encapsulate the
// os.Process. The wrapper also includes channels for errors and closing up shop when gracefully shutting down.
type Proc struct {
	id string

	process   *os.Process
	procState *os.ProcessState
	procAttr  *os.ProcAttr

	finished bool
	errch    chan error
	waitDone chan struct{}
}

type Job struct {
	index 	  int
	proc      Proc
	
	priority  JobPriority
	timeSlice time.Time
	state     JobState
}

type JobState string

// JobPriority separates the four job queues from one another, giving each queue a certain "weight" or
// priority on a scale of 0-4, with 0 being of the highest priority, and 4 as the last.
type JobPriority uint32

// MinPriority
func (p JobPriority) MinPriority() int { return 0 }

// MaxPriority
func (p JobPriority) MaxPriority() int { return 0 }
