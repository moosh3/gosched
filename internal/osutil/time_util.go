package gomlfq

import "time"

func CPULoop(q Queue, p Proc) {
	min := q.timeSlice
	time.Sleep(min)
	for {
		t0 := time.Now()
		run(p)
		// duration of run
		diff := time.Now().Sub(t0)
		if diff < min {
			delay := min - diff
			// ensures the proc uses the
			// full timeSlice
			time.Sleep(delay)
		}
	}
}
