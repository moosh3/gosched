// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gosched

type Scheduler struct{}

// JobScheduler implements a CPU process scheduler, handing incoming Jobs and
// manipulating related priority queues accordingly.
type JobScheduler interface {
	Promote(Job, priority int) error
	PromoteAll([]Job)
	Demote(Job, priority int)
	Recv(jobChan chan struct{})
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}
