package task

import (
	"log"
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
	model, _ := getAddTaskModel(c)
	task := h.usercase.AddNew(model)
	c.JSON(200, task)
}

func (h *HttpTaskHandler) getById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	log.Println("HEEEEEEEEEEEEEEEEEEEE", id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	task := h.usercase.GetOne(int(id))
	c.JSON(200, task)
}
