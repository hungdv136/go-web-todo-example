package task

import (
	"net/http"
	"strconv"

	"todo/common"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

func (r *TaskRequest) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&r.UserId: binding.Field{
			Form:         "userId",
			Required:     true,
			ErrorMessage: "userId is a required field.",
		},
		&r.Title:   "title",
		&r.Content: "content",
	}
}

func getAddTaskModel(c *gin.Context) TaskRequest {
	var model TaskRequest
	common.AbortWithStatusIfError(c, common.Bind(c, &model), http.StatusBadRequest)
	return model
}

func getRequestTaskId(c *gin.Context) int {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	common.AbortWithStatusIfError(c, err, http.StatusBadRequest)
	return int(id)
}
