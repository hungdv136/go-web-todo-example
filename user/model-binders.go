package user

import (
	"net/http"

	"todo/common"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

func (r *LoginModel) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&r.Username: binding.Field{
			Form:         "useranme",
			Required:     true,
			ErrorMessage: "useranme is a required field.",
		},
		&r.Password: "password",
	}
}

func getLoginModel(c *gin.Context) LoginModel {
	var model LoginModel
	common.Bind(c, &model)
	return model
}
