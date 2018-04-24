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

// GoScheduler implements a CPU process scheduler, handing incoming Jobs and
// manipulating related priority queues accordingly.
type GoScheduler interface {
	SocketController
	QueueController
	Listener

	Promote(ctx context.Context, j Job, pr int) error       // Promote a job to a new (existing) queue
	PromoteAll(ctx context.Context, oldPr, newPr int) error // Promote all jobs from one exisiting queue to another existing queue
	Demote(ctx context.Context, j Job, pr int) error        // Demote a job to a new (existing) queue
	DemoteAll(ctx context.Context, oldPr, newPr int) error  // Demote all jobs from one exisiting queue to another existing queue

	Start() error
	Stop()
}

// SocketController supplies an interface for interacting with TCP sockets, parsing a stream
// of incoming Job payloads from a specific TCP socket
type SocketController interface {
	Recv(jobChan chan struct{}) error
	Dial(socket string) error
	Close(ctx context.Context) error
	Error() error
}

// QueueController is sort of policy specific - but overall, it can be considered the controller for the
// queues being used via any policy. Priority queues are expected, and identified via a mapped int to the string representation.
type QueueController interface {
	Start(pr int) error
	Stop(pr int)
	Enqueue(pr int, item interface{}) error
	Dequeue(pr int, item interface{}) error
}

// Listener provides an interface that can act as a gateway between any activity from an external source
// (i.e. a JSON payload from a TCP socket). It's basically a buffer between SocketController and GoScheduler,
// to better handle queue scheduling.
type Listener interface {
}
