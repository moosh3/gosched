// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gosched

import (
	"context"
)

// Scheduler contains information needed for use when implementing
// the JobScheduler interface for a given scheduling policy
type Scheduler struct {
	hMu    synxx.Mutex // protects hot
	policy string
}

// JobScheduler implements a CPU process scheduler, handing incoming Jobs and
// manipulating related priority queues accordingly.
type JobScheduler interface {
	Promote(context.Context, Job) error
	PromoteAll(context.Context, []Job) error
	Demote(context.Context) error
	Recv(jobChan chan struct{})
	Close(context.Context) error
	Error() error
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) run() error {
	// ...

	go func() {
		for {
			s.cond.L.Lock()

			isWaiting = true
			for len(a.slots) <= 0 && (ctx.Err() == nil) {
				a.cond.Wait()
			}
			isWaiting = false

			if ctx.Err() != nil {
				a.cond.L.Unlock()
				return
			}

			item := s.queueList[len(s.queueList)-1]
			a.cond.L.Unlock()

			select {
			case output <- item: // good, dequeued
			case <-item.trigger: // ejected
			case <-ctx.Done(): // timeout or cancel from caller
			}
		}
	}()
	return output
}
