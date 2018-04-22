// Package gosched is a set of packages that provide tools for constructing a Job scheduler
//
// gosched contains the following packages:
//
// The channel package contains helper functions for working with channels and goroutines.
//
// The constants package provides definitions for various, well, constants
//
// The context package allows for context switching between jobs given a scheduling policy
//
// The policy package is one of the more important packages, supplying scheduling policies
// for FIFO queue, Round Robin, and Multi-level Feedback queues
//
// The signal package provides a set of signals used to trigger proccesses and other system calls
//
// Example Usage
//
// The following is a complete example using gosched to create a job scheduler:
//    import (
//        "github.com/aleccunningham/gosched"
//    )
//
//    func main() {
//        ctx := context.Background()
//        done := make(chan bool)
//        // ...
//        scheduler := gosched.NewScheduler("mlfq")
//        scheduler.Begin()
//
// Policies
//
// Policies define exactly how the scheduler will interact, commonly called a scheduling algorithm
// or scheduling policy. gosched provides SJF (shortest job first), FIFO (first come first server),
// Round Robin, and MLFQ (Multi-level feedback queue). Note that the MLFQ policy uses both FIFO and Round Robin,
// and is the only policy provided that will not return terrible turnaround or response time.
package gosched

import (
	// channel package
	_ "github.com/aleccunningham/gosched/channel"
	// queue package
	_ "github.com/aleccunningham/gosched/queue"
	// signal package
	_ "github.com/aleccunningham/gosched/signal"
)
