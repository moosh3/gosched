package gosched

// Dispatcher is the component of the scheduler that handles the
// mechanism of actually getting that process to run on the processor.
// This requires loading the saved context of the selected Job, which is
// stored in the process controll block and comprimises the set of
// registers, stack pointer, flags, and a pointer to the memory mapping (typicall a pointer to the page table).
//
// Once this context is loaded, the dispatcher switches to user mode via a
// return from interrupt operation that causes the job to execute from the location
// that was saved on the stack at the time the program stopped running, either via an
// interrupt or a system call.
type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}
