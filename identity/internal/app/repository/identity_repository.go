package repository

import (
	"github.com/google/uuid"
	"github.com/hoangkhoachau/MicroserviceShop/identity/internal/app/model"
)

type IIdentityRepository interface {
	AddUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUserById(id int) error
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserRolesByUserId(id int) ([]string, error)
	CheckTokenExist(uuid string) bool
	AddToken(tokenDetail *model.TokenDetail) error
	DeleteTokenByUUID(uuid uuid.UUID) error
	VerifyUserEmailByCode(code string) error
}
