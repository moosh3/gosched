package gosched

import "sync"

// IO represents CPU locks when a process requests IO usage
type IO struct {
	sync.Mutex
}

// CPU represents the scheduler's CPU on which it will schedule jobs onto
type CPU struct {
	sync.Mutex
}
