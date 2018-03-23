package main

import (
	"todo/middleware"

	"github.com/gin-gonic/gin"

	task "todo/task"
	user "todo/user"
)

func main() {
	r := gin.Default()

	taskRepository := task.NewSqlRepository()
	taskHandler := task.NewTaskHandler(taskRepository)
	userHandler := user.NewUserHandler()

	userHandler.Build(r)
	r.Use(middleware.Authentication())
	{
		taskHandler.Build(r)
	}

	r.Run(":4200")
}
