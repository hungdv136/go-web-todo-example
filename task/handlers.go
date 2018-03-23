package task

type TaskHandler struct {
	repository TaskRepository
}

func NewTaskHandler(repository TaskRepository) *TaskHandler {
	return &TaskHandler{repository}
}

func (h *TaskHandler) addNew(task TaskRequest) *Task {
	return h.repository.Add(task.UserId, task.Title, task.Content)
}

func (h *TaskHandler) getAll() {

}

func (h *TaskHandler) complete(taskId int) {

}

func (h *TaskHandler) cancel(taskId int) {

}
