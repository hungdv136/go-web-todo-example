package task

import (
	"net/http"
	"todo/common"
	"todo/middleware"

	"github.com/gin-gonic/gin"
)

type HttpTaskHandler struct {
	usercase TaskUsercase
}

func NewHttpTaskHandler(usercase TaskUsercase) HttpTaskHandler {
	return HttpTaskHandler{usercase}
}

func (h *HttpTaskHandler) Build(engine *gin.Engine) {
	authMiddleware := middleware.Authentication()
	engine.POST("task/add", authMiddleware, h.addTaskHandler)
	engine.GET("task/:id", authMiddleware, h.getById)
	engine.POST("task/complete/:id", authMiddleware, h.complete)
}

func (h *HttpTaskHandler) addTaskHandler(c *gin.Context) {
	model := getAddTaskModel(c)
	task, err := h.usercase.AddNew(model)
	common.AbortWithStatusIfError(c, err, http.StatusInternalServerError)
	c.JSON(200, task)
}

func (h *HttpTaskHandler) getById(c *gin.Context) {
	id := getRequestTaskId(c)
	task, err := h.usercase.GetOne(id)
	common.AbortWithStatusIfError(c, err, http.StatusInternalServerError)
	c.JSON(200, task)
}

func (h *HttpTaskHandler) complete(c *gin.Context) {
	id := getRequestTaskId(c)
	err := h.usercase.Complete(int(id))
	common.AbortWithStatusIfError(c, err, http.StatusInternalServerError)
}
