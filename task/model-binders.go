package task

import (
	"net/http"

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

func getAddTaskModel(c *gin.Context) (TaskRequest, error) {
	var model TaskRequest
	err := common.Bind(c, &model)
	return model, err
}
