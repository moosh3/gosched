// Copyright 2018 Alec Cunningham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package policy provides scheduling policies (policies == algorithms)

fifo_queue provides a queue interface that can be used as a Job queue
with the use of a Job Scheduler

mlf_queue is a JobScheduler implementation, scheduling Jobs on a virtualized CPU

rrqueue provides a container for priority queues and a simple round-robin scheduled consumer.

*/
package policy
