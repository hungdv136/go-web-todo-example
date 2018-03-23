package middleware

import (
	"net/http"
	"strings"

	"todo/common"

	"github.com/gin-gonic/gin"
)

const (
	headerName = "Authorization"
	jwtType    = "Bearer "
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(headerName)

		if !strings.HasPrefix(token, jwtType) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenRaw := strings.Split(token, jwtType)[1]
		valid := common.ValidateToken(tokenRaw)

		if !valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
