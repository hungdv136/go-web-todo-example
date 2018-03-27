package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

func Bind(c *gin.Context, userStruct binding.FieldMapper) {
	err := binding.Bind(c.Request, userStruct)
	AbortWithStatusIfError(c, err, http.StatusBadRequest)
}
