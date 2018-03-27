package task

type TaskUsercase struct {
	repository TaskRepository
}

func NewTaskUsercase(repository TaskRepository) TaskUsercase {
	return TaskUsercase{repository}
}

func (u *TaskUsercase) AddNew(task TaskRequest) *Task {
	newTask := Task{Title: task.Title, Content: task.Content, UserId: task.UserId}
	newTask.Id = u.repository.Add(newTask)
	return &newTask
}

func (u *TaskUsercase) Fecth() []Task {
	return u.repository.Fetch()
}

func (u *TaskUsercase) Complete(taskId int) {
	u.update(taskId, func(task *Task) {
		task.Complete()
	})
}

func (u *TaskUsercase) Cancel(taskId int) {
	u.update(taskId, func(task *Task) {
		task.Cancel()
	})
}

func (u *TaskUsercase) GetOne(taskId int) Task {
	return u.repository.GetById(taskId)
}

func (u *TaskUsercase) update(taskId int, updator func(task *Task)) {
	task := u.repository.GetById(taskId)
	updator(&task)
	u.repository.Update(task)
}
