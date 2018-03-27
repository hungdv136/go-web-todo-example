package task

type TaskUsercase struct {
	repository TaskRepository
}

func NewTaskUsercase(repository TaskRepository) TaskUsercase {
	return TaskUsercase{repository}
}

func (u *TaskUsercase) AddNew(task TaskRequest) (*Task, error) {
	newTask := Task{Title: task.Title, Content: task.Content, UserId: task.UserId}
	taskId, err := u.repository.Add(newTask)
	newTask.Id = taskId
	return &newTask, err
}

func (u *TaskUsercase) Fecth() ([]Task, error) {
	tasks, err := u.repository.Fetch()
	return tasks, err
}

func (u *TaskUsercase) Complete(taskId int) {
	u.update(taskId, func(task *Task) { task.Complete() })
}

func (u *TaskUsercase) Cancel(taskId int) {
	u.update(taskId, func(task *Task) { task.Cancel() })
}

func (u *TaskUsercase) GetOne(taskId int) (Task, error) {
	task, err := u.repository.GetById(taskId)
	return task, err
}

func (u *TaskUsercase) update(taskId int, updator func(task *Task)) error {
	task, _ := u.repository.GetById(taskId)
	updator(&task)
	return u.repository.Update(task)
}
