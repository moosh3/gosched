package gomlfq

import (
	"container/heap"
	// "github.com/aleccunningham/gomlfq/sched"
)

// Heap implements heap.Interface and holds Items, managing them
// via the use of a in-memory FIGO queue
type Heap []*Job

func (pq Heap) Len() int { return len(pq) }

func (pq Heap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].jid = i
	pq[j].jid = j
}

func (pq *Heap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Job)
	item.jid = n
	*pq = append(*pq, item)
}

func (pq *Heap) Pop() interface{} {
	old := *pq
	n := len(old)
	job := old[n-1]
	job.jid = -1 // for safety
	*pq = old[0 : n-1]
	return job
}

// update modifies the priority and value of an Item in the queue.
func (pq *Heap) update(job *Job, value string, priority int) {
	job.value = value
	job.priority = priority
	heap.Fix(pq, job.jid)
}
