package main

import (
	"todo/infrastructure"

	"github.com/gin-gonic/gin"

	task "todo/task"
	user "todo/user"
)

func main() {
	r := gin.Default()
	db := infrastructure.Init()
	defer db.Close()

	taskRepository := task.NewSqlRepository()
	taskUsercase := task.NewTaskUsercase(taskRepository)
	taskHandler := task.NewHttpTaskHandler(taskUsercase)
	identityUsercase := user.NewIdentityUsercase()
	userHandler := user.NewHttpIdentityHandler(identityUsercase)

	userHandler.Build(r)
	taskHandler.Build(r)

	r.Run(":4200")
}
