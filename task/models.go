package task

type TaskRequest struct {
	UserId  int
	Title   string
	Content string
}

type Status int8

const (
	Pending   = Status(0)
	Started   = Status(1)
	Completed = Status(2)
	Canceled  = Status(3)
)

type Task struct {
	Id      int    `gorm:"column:Id;AUTO_INCREMENT;primary_key"`
	UserId  int    `gorm:"column:UserId"`
	Title   string `gorm:"column:Title;type:varchar(250)"`
	Content string `gorm:"column:Content"`
	Status  Status `gorm:"column:Status"`
}

func (Task) TableName() string {
	return "task"
}

func CreateTask(userId int, title string, content string) *Task {
	return &Task{
		UserId:  userId,
		Title:   title,
		Content: content,
		Status:  Pending,
	}
}

func (item *Task) Complete() {
	item.Status = Completed
}

func (item *Task) Start() {
	item.Status = Started
}

func (item *Task) Cancel() {
	item.Status = Canceled
}
