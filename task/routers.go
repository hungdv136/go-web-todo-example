package task

import (
	"github.com/gin-gonic/gin"
)

func (h *TaskHandler) Build(engine *gin.Engine) {
	engine.POST("task/add/", func(c *gin.Context) {
		model, _ := getAddTaskModel(c)
		task := h.addNew(model)
		c.JSON(200, task)
	})
}
