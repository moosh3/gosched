// Package channel contains methods to manipulate, read, and write via the channel primitive
package channel

import "sync"

func generateStream(done <-chan interface{}, jobs ...interface{}) <-chan interface{} {
	jobStream := make(chan interface{})
	go func() {
		defer close(jobStream)
		for _, i := range jobs {
			select {
			case <-done:
				return
			case jobStream <- i:
			}
		}
	}()
	return jobStream
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
