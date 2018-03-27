package task

import (
	"todo/infrastructure"

	"github.com/jinzhu/gorm"
)

type TaskRepository interface {
	GetById(taskId int) Task
	Fetch() []Task
	Add(task Task) int
	Update(task Task)
}

type sqlTaskRepository struct {
	db *gorm.DB
}

func NewSqlRepository() TaskRepository {
	return &sqlTaskRepository{infrastructure.GetDB()}
}

func (r *sqlTaskRepository) Add(task Task) int {
	if err := r.db.Create(&task).Error; err != nil {
		panic(err)
	}
	return task.Id
}

func (r *sqlTaskRepository) Update(task Task) {
	if err := r.db.Update(task).Error; err != nil {
		panic(err)
	}
}

func (r *sqlTaskRepository) GetById(taskId int) Task {
	var task Task
	r.db.First(&task, taskId)
	// if err := r.db.First(&task, taskId); err != nil {
	// 	log.Panicln("ERRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR", err.Error, task)
	// 	panic(err)
	// }
	return task
}

func (r *sqlTaskRepository) Fetch() []Task {
	var tasks []Task
	if err := r.db.Find(tasks); err != nil {
		panic(err)
	}
	return make([]Task, 5)
}
