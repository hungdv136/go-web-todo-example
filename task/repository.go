package task

type TaskRepository interface {
	Add(userId int, title string, content string) *Task
	Complete(task *Task)
	Cancel(task *Task)
	Start(task *Task)
}

type sqlTaskRepository struct{}

func NewSqlRepository() TaskRepository {
	return &sqlTaskRepository{}
}

func (repository *sqlTaskRepository) Add(userId int, title string, content string) *Task {
	return CreateTask(userId, title, content)
}

func (repository *sqlTaskRepository) Complete(task *Task) {
	task.Complete()
}

func (repository *sqlTaskRepository) Cancel(task *Task) {
	task.Cancel()
}

func (repository *sqlTaskRepository) Start(task *Task) {
	task.Start()
}
