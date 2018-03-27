package common

import (
	"github.com/gin-gonic/gin"
)

func AbortWithStatusIfError(c *gin.Context, err error, code int) {
	if err != nil {
		c.AbortWithStatus(code)
		panic(err)
	}
}
