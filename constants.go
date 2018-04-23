// Package constants provides values for scheduler policies
package gosched

// Process switch behavior
const (
	SwitchOnIO  = "SCHED_SWITCH_ON_IO"
	SwitchOnEnd = "SCHED_SWITCH_ON_END"
)

// IO finished behavior
const (
	IORunLater     = "IO_RUN_LATER"
	IORunImmediate = "IO_RUN_IMMEDIATE"
)

// Process states
const (
	Running = "STATE_RUNNING"
	Ready   = "STATE_READY"
	Done    = "DONE"
	Wait    = "WAITING"
)

// Members of the process struct
const (
	code_       = "PROC_CODE"
	pc_         = "PROC_PC"
	pid_        = "PROC_ID"
	proc_state_ = "PROC_STATE"
)

// Process options
const (
	DoCompute = "cpu"
	DoIO      = "io"
)
