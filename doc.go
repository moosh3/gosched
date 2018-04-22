// Package gosched is a set of packages that provide tools for constructing a Job scheduler
//
// gosched contains the following packages:
//
// The channel package contains helper functions for working with channels and goroutines.
//
// The queue package provides a FIFO queue in-memory implementation, along with a round robin queue
//
// The signal package provides a set of signals used to trigger proccesses and other system calls

package gosched

import (
	// channel package
	_ "github.com/aleccunningham/gosched/channel"
	// queue package
	_ "github.com/aleccunningham/gosched/queue"
	// signal package
	_ "github.com/aleccunningham/gosched/signal"
)
