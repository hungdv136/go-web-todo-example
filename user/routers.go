package user

import (
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Build(engine *gin.Engine) {
	engine.POST("user/login/", func(c *gin.Context) {
		model, _ := getLoginModel(c)
		token, err := h.login(model)
		if err != nil {
			c.Error(err)
		} else {
			c.JSON(200, token)
		}
	})
}
