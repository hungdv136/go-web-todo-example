package user

import (
	"todo/common"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) login(model LoginModel) (string, error) {
	return common.NewWithClaims(model.Username)
}
