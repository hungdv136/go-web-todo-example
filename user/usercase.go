package user

import (
	"todo/common"
)

type IdentityUsercase struct {
}

func NewIdentityUsercase() IdentityUsercase {
	return IdentityUsercase{}
}

func (h *IdentityUsercase) login(model LoginModel) (string, error) {
	return common.NewWithClaims(model.Username)
}
