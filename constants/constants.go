// Package constants provides values for scheduler policies
package constants

const (
	// Process switch behavior
	SwitchOnIO  = "SCHED_SWITCH_ON_IO"
	SwitchOnEnd = "SCHED_SWITCH_ON_END"

	// IO finished behavior
	IORunLater     = "IO_RUN_LATER"
	IORunImmediate = "IO_RUN_IMMEDIATE"

	// Process states
	Running = "STATE_RUNNING"
	Ready   = "STATE_READY"
	Done    = "DONE"
	Wait    = "WAITING"

	// Members of the process struct
	code_       = "PROC_CODE"
	pc_         = "PROC_PC"
	pid_        = "PROC_ID"
	proc_state_ = "PROC_STATE"

	// Process options
	DoCompute = "cpu"
	DoIO      = "io"
)
