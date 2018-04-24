package policy

import "sync"

type sjf struct {
	sync.Mutex
}

func (s *sjf) run() error {
	// ...

	go func() {
		for {
			s.cond.L.Lock()

			isWaiting = true
			for len(a.slots) <= 0 && (ctx.Err() == nil) {
				a.cond.Wait()
			}
			isWaiting = false

			if ctx.Err() != nil {
				a.cond.L.Unlock()
				return
			}

			item := s.queueList[len(s.queueList)-1]
			a.cond.L.Unlock()

			select {
			case output <- item: // good, dequeued
			case <-item.trigger: // ejected
			case <-ctx.Done(): // timeout or cancel from caller
			}
		}
	}()
	return output
}
