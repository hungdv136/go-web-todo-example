package user

import (
	"github.com/gin-gonic/gin"
)

type HttpIdentityHandler struct {
	usercase IdentityUsercase
}

func NewHttpIdentityHandler(usercase IdentityUsercase) HttpIdentityHandler {
	return HttpIdentityHandler{usercase}
}

func (h *HttpIdentityHandler) Build(engine *gin.Engine) {
	engine.POST("user/login/", h.loginHandler)
}

func (h *HttpIdentityHandler) loginHandler(c *gin.Context) {
	model, _ := getLoginModel(c)
	token, err := h.usercase.login(model)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(200, token)
	}
}
