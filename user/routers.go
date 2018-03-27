package user

import (
	"net/http"
	"todo/common"

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
	common.AbortWithStatusIfError(c, err, http.StatusBadRequest)
	token, err := h.usercase.login(model)
	common.AbortWithStatusIfError(c, err, http.StatusInternalServerError)
	c.JSON(200, token)
}
