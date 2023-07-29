package main

// LoadBalancer represents the load balancer that distributes tasks among workers.
type LoadBalancer struct {
	workers []*Worker
	next    int // Index of the next worker to receive a task
}

// NewLoadBalancer creates a new LoadBalancer instance.
func NewLoadBalancer(workers []*Worker) *LoadBalancer {
	return &LoadBalancer{
		workers: workers,
	}
}

// GetNextWorker returns the next worker to receive a task using round-robin scheduling.
func (lb *LoadBalancer) GetNextWorker() *Worker {
	worker := lb.workers[lb.next]
	lb.next = (lb.next + 1) % len(lb.workers)
	return worker
}