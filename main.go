package main

import (
	"todo/infrastructure/cache"
	"todo/infrastructure/database"

	"github.com/gin-gonic/gin"

	task "todo/task"
	user "todo/user"
)

func main() {
	engine := gin.Default()
	db := database.Init()
	cache := cache.Init()

	defer func() {
		db.Close()
		cache.Empty()
	}()

	taskRepository := task.NewSqlRepository(db)
	taskUsercase := task.NewTaskUsercase(taskRepository)
	taskHandler := task.NewHttpTaskHandler(taskUsercase)
	identityUsercase := user.NewIdentityUsercase()
	userHandler := user.NewHttpIdentityHandler(identityUsercase)

	userHandler.Build(engine)
	taskHandler.Build(engine)

	engine.Run(":4200")
}
