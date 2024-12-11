package usecase

import "github.com/hoangkhoachau/MicroserviceShop/identity/internal/app/dto"

type IIdentityService interface {
	CreateUser(createUserRequest dto.CreateUserRequest) error
	LoginUser(loginUserRequest dto.LoginUserRequest) (*dto.TokenResponse, error)
	VerifyUserByCode(code string) error
	CheckTokenExist(userId int) bool
}
