// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package gosched contains a set of libraries providing tools for constructing a Job scheduler, based on a configuration of Job schemas and scheduling policy. A typical use is scheduling Jobs onto one or more queues based on associated identifiers (timeSlice, jobPriority). The main interface, the Scheduler, can use any scheduling policy that implements its interface. Currently that includes shortest-job-first, fifo, and round robin. Policies can be found in the gosched/policy package.

Package gosched/channel

contains helper functions for working with channels and goroutines.

Package gosched/constants

provides definitions for various, well, constants

Package gosched/context

allows for context switching between jobs given a scheduling policy

Package gosched/policy

is one of the more important packages, supplying scheduling policies for FIFO queue, Round Robin, and Multi-level Feedback queues

Package gosched/signal

provides a set of signals used to trigger proccesses and other system calls

Example Usage

The following is a complete example using gosched to create a job scheduler:
  import (
      "github.com/aleccunningham/gosched"
  )

  func main() {
      ctx := context.Background()
      done := make(chan bool)
       ...
      scheduler := gosched.NewScheduler("mlfq")
      scheduler.Begin()

Policies

Policies define exactly how the scheduler will interact, commonly called a scheduling algorithm
or scheduling policy. gosched provides SJF (shortest job first), FIFO (first come first server),
Round Robin, and MLFQ (Multi-level feedback queue). Note that the MLFQ policy uses both FIFO and Round Robin,
and is the only policy provided that will not return terrible turnaround or response time.
*/
package gosched

import (
	// channel package
	_ "github.com/aleccunningham/gosched/channel"
	// constants package
	_ "github.com/aleccunningham/gosched/constants"
	// context package
	_ "github.com/aleccunningham/gosched/context"
	// policy package
	_ "github.com/aleccunningham/gosched/policy"
	// signal package
	_ "github.com/aleccunningham/gosched/signal"
)
