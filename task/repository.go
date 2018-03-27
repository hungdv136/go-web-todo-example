package task

import (
	"todo/infrastructure"

	"github.com/jinzhu/gorm"
)

type TaskRepository interface {
	GetById(taskId int) (Task, error)
	Fetch() ([]Task, error)
	Add(task Task) (int, error)
	Update(task Task) error
}

type sqlTaskRepository struct {
	db *gorm.DB
}

func NewSqlRepository() TaskRepository {
	return &sqlTaskRepository{infrastructure.GetDB()}
}

func (r *sqlTaskRepository) Add(task Task) (int, error) {
	err := r.db.Create(&task).Error
	return task.Id, err
}

func (r *sqlTaskRepository) Update(task Task) error {
	return r.db.Update(task).Error
}

func (r *sqlTaskRepository) GetById(taskId int) (Task, error) {
	var task Task
	err := r.db.First(&task, taskId).Error
	return task, err
}

func (r *sqlTaskRepository) Fetch() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(tasks).Error
	return make([]Task, 5), err
}
