package gomlfq

import "sync"

func generateStream(done <-chan interface{}, procs ...Proc) <-chan Proc {
	procStream := make(chan Proc)
	go func() {
		defer close(procStream)
		for _, i := range procs {
			select {
			case <-done:
				return
			case procStream <- i:
			}
		}
	}()
	return procStream
}

func fanOut(workers int) {
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(i)
	}
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	// when passed a channel, multiplex will read from
	// the channnel and pass the read value onto
	// the multiplexedStream channel
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
