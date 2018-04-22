package gosched

// "github.com/aleccunningham/gomlfq/queue"

type Scheduler struct{}

// JobScheduler implements a CPU process scheduler, handing incoming Jobs and
// manipulating related priority queues accordingly.
type JobScheduler interface {
	Promote(Job, priority int) error
	PromoteAll([]Job)
	Demote(Job, priority int)
	Recv(jobChan chan struct{})
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}
