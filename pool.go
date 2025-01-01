package worker

import (
	"sync"
)

type Pool struct {
	size    int
	jobChan chan Job
	wg      sync.WaitGroup
}

func (p *Pool) worker() {
	defer p.wg.Done()

	for job := range p.jobChan {
		job.Process()
	}
}

func (p *Pool) Start() {
	p.jobChan = make(chan Job)

	p.wg.Add(p.size)
	for i := 0; i < p.size; i++ {
		go p.worker()
	}
}

func (p *Pool) AddJob(j Job) {
	p.jobChan <- j
}

func (p *Pool) Close() {
	close(p.jobChan)
	p.wg.Wait()
}
