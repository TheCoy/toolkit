package routinepool

//Pool defines the pool for goroutines
type Pool struct {
	EntryChannel chan *Task
	JobChannel   chan *Task
	poolSize     int
}

func NewPool(size int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		JobChannel:   make(chan *Task),
		poolSize:     size,
	}
}

func (p *Pool) worker(id int) {
	for task := range p.JobChannel {
		task.Execute()
	}
}

func (p *Pool) Run() {
	for index := 0; index < p.poolSize; index++ {
		go p.worker(index)
	}

	for job := range p.EntryChannel {
		p.JobChannel <- job
	}
}
