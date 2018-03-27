package main

import (
	"todo/infrastructure"
	"todo/middleware"

	"github.com/gin-gonic/gin"

	task "todo/task"
	user "todo/user"
)

func main() {
	r := gin.Default()
	infrastructure.Init()

	taskRepository := task.NewSqlRepository()
	taskUsercase := task.NewTaskUsercase(taskRepository)
	taskHandler := task.NewHttpTaskHandler(taskUsercase)
	identityUsercase := user.NewIdentityUsercase()
	userHandler := user.NewHttpIdentityHandler(identityUsercase)

	userHandler.Build(r)
	r.Use(middleware.Authentication())
	{
		taskHandler.Build(r)
	}

	r.Run(":4200")
}
