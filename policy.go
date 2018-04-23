// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gosched

const (
	Other Policy = 0
	FIFO  Policy = 1
	RR    Policy = 2
	Batch Policy = 3
	Idle  Policy = 4
)

func (p Policy) String() string {
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

type Policy interface {
	Enqueue(int, interface{}) error
	Dequeue(int) (interface{}, error)
}
