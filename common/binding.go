package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

func Bind(c *gin.Context, userStruct binding.FieldMapper) error {
	err := binding.Bind(c.Request, userStruct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return err
	}
	return err
}
