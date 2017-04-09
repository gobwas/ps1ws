type Pool struct {
	work chan func()
	sem  chan struct{}
}

// START OMIT
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan size),
	}
}

func (p *Pool) Schedule(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	for {
		task()
		task = <-p.work
	}
}

// END OMIT
