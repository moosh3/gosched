// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gosched

import (
	"os"
	"time"
)

// Job is a wrapper/abstraction for a traditional os.Process call by giving it
// attributes and methods that allow it to be handled and executed by a JobScheduler with any
// implemented JobPolicy. Note that once a Job is scheduled, it runs its Job.process.
type Job struct {
	jid     int         `json:"jid"`     // jid is the position in a given slice of Jobs
	pid     int         `json:"pid"`     // pid of the Jobs associated process
	process *os.Process `json:"process"` // proc is the os process that the job is scheduled to run

	priority  JobPriority  `json:"jobPriority"` // corresponds to the priority queues in order to place the job on the right queue
	timeSlice time.Time    `json:"timeSlice"`   // time io/mem blocks of which the scheduler must handle
	state     JobStateType `json:"state"`
}

// JobPriority separates the four job queues from one another, giving each queue a certain "weight" or
// priority on a scale of 0-4, with 0 being of the highest priority, and 4 as the last.
type JobPriority int

// MinPriority returns the min priority provided by the scheduler, which is the same
// value as the priority on the lowest priority FIFO queue
func (p JobPriority) MinPriority() int { return 0 }

// MaxPriority returns the max priority provided by the scheduler, which is the same
// value as the priority on the highest priority FIFO queue
func (p JobPriority) MaxPriority() int { return 0 }

// JobStateType is an integer representation of the current state a specific Job
type JobStateType int

// Possible states a Job could be in at any given time
const (
	JobStateNone    JobStateType = iota // unititialized
	JobStateReady                       // job is ready to be scheduled
	JobStateWait                        // job request is waiting
	JobStateRunning                     // job request is running its process
	JobStateDone                        // job request has finished
)

// MoveToReady changes a Jobs state to Ready
func (j *Job) MoveToReady(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Ready"
	job.process = proc
}

// MoveToWait changes a Jobs state to Ready
func (j *Job) MoveToWait(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Wait"
	job.process = proc
}

// MoveToRunning changes a Jobs state to Ready
func (j *Job) MoveToRunning(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Running"
	job.process = proc
}

// MoveToDone changes a Jobs state to Ready
func (j *Job) MoveToDone(pid int) {
	proc := findByPid(pid)

	var job Job
	job.state = "Done"
	job.process = proc
}

func findByJid(jid int) *os.Process { return nil }

func findByPid(pid int) *os.Process { return nil }
