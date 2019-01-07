package dispatcher

import (
	"runtime"
)

const Concurrency = 32

type Dispatcher struct {
	WorkerPool chan chan Job
	jobQueue   chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
		jobQueue:   make(chan Job, runtime.NumCPU()*(Concurrency)),
		maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) AddJob(job Job) {
	d.jobQueue <- job
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}()
		}
	}
}
