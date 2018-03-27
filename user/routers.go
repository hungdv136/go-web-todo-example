package user

import (
	"net/http"

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
	model, err := getLoginModel(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	token, err := h.usercase.login(model)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(200, token)
	}
}
