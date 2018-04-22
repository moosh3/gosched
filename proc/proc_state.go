package proc

import (
	"os"
	"syscall"
)

type ProcessState struct {
	pid    int
	status syscall.WaitStatus
	rusage *syscall.Rusage
}

func (p *ProcessState) Exited() bool          { return true }
func (p *ProcessState) Pid() int              { return 0 }
func (p *ProcessState) String() string        { return " " }
func (p *ProcessState) Success() bool         { return true }
func (p *ProcessState) Sys() interface{}      { return nil }
func (p *ProcessState) SysUsage() interface{} { return nil }

func MoveToReady(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Ready"
	job.process = proc
}

func MoveToWait(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Wait"
	job.process = proc
}

func MoveToRunning(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Running"
	job.process = proc
}

func MoveToDone(pid int) {
	proc := findByPid(pid)

	var job Proc
	job.state = "Done"
	job.process = proc
}

func findByPid(pid int) *os.Process { return nil }
