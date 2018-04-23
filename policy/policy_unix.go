// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package policy

import (
	"syscall"
	"unsafe"
)

const (
	Other UnixPolicy = iota
	FIFO
	RR
	Batch
	Idle
)

func (p UnixPolicy) String() string {
	switch p {
	case Other:
		return "Other"
	case FIFO:
		return "FIFO"
	case RR:
		return "RR"
	case Batch:
		return "Batch"
	case Idle:
		return "Idle"
	}
	return "Unknown"
}

type UnixPolicy uintptr

type Param struct{ Priority int }

func (p Policy) SetPolicy(pid int, policy UnixPolicy, param *Param) error {
	_, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_SETSCHEDULER,
		uintptr(pid),
		uintptr(policy),
		uintptr(unsafe.Pointer(param)),
	)
	if e != 0 {
		return e
	}
	return nil
}

func (p Policy) GetPolicy(pid int) (UnixPolicy, error) {
	r0, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_GETSCHEDULER, uintptr(pid), 0, 0,
	)
	if e != 0 {
		return 0, e
	}
	return Policy(r0), nil
}

func SetParam(pid int, param *Param) error {
	_, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_SETPARAM,
		uintptr(pid),
		uintptr(unsafe.Pointer(param)),
		0,
	)
	if e != 0 {
		return e
	}
	return nil
}

func GetParam(pid int, param *Param) error {
	_, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_GETPARAM,
		uintptr(pid),
		uintptr(unsafe.Pointer(param)),
		0,
	)
	if e != 0 {
		return e
	}
	return nil
}
