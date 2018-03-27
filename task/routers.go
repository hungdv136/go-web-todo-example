package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpTaskHandler struct {
	usercase TaskUsercase
}

func NewHttpTaskHandler(usercase TaskUsercase) HttpTaskHandler {
	return HttpTaskHandler{usercase}
}

func (h *HttpTaskHandler) Build(engine *gin.Engine) {
	engine.POST("task/add", h.addTaskHandler)
	engine.GET("task/:id", h.getById)
}

func (h *HttpTaskHandler) addTaskHandler(c *gin.Context) {
	model, err := getAddTaskModel(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	task, err := h.usercase.AddNew(model)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(200, task)
}

func (h *HttpTaskHandler) getById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	task, err := h.usercase.GetOne(int(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(200, task)
}
