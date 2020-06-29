package routinepool

//Task defines task
type Task struct {
	f func() error
}

func NewTask(arg_f func() error) *Task {
	return &Task{
		f: arg_f,
	}
}

func (task *Task) Execute() {
	task.f()
}
