package worker

import (
	"golang.org/x/sync/errgroup"
)

type Pool struct {
	size    int
	jobChan chan Job
	g       *errgroup.Group
}

func NewPool(poolSize int) Pool {
	return Pool{
		size:    poolSize,
		jobChan: make(chan Job),
		g:       &errgroup.Group{},
	}
}

func (p *Pool) AddJob(j Job) {
	p.jobChan <- j
}

func (p *Pool) Close() error {
	close(p.jobChan)
	return p.g.Wait()
}

func (p *Pool) Start() {
	p.g.SetLimit(p.size)

	for i := 0; i < p.size; i++ {
		p.g.Go(p.worker)
	}
}

func (p *Pool) worker() error {

	for job := range p.jobChan {
		if err := job.Process(); err != nil {
			return err
		}
	}

	return nil
}
