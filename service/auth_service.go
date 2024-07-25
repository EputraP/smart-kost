package service

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
	"smart-kost-backend/util/hasher"
	"smart-kost-backend/util/tokenprovider"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(input dto.LoginBody) (*dto.LoginResponse, error)
	SignUp(input dto.User) (*dto.User, error)
}

type authService struct {
	userRepo    repository.UserRepository
	hasher      hasher.Hasher
	jwtProvider tokenprovider.JWTTokenProvider
}

type AuthServiceConfig struct {
	UserRepo    repository.UserRepository
	Hasher      hasher.Hasher
	JwtProvider tokenprovider.JWTTokenProvider
}

func NewAuthService(config AuthServiceConfig) AuthService {
	return &authService{
		userRepo:    config.UserRepo,
		hasher:      config.Hasher,
		jwtProvider: config.JwtProvider,
	}
}

func (ts authService) Login(input dto.LoginBody) (*dto.LoginResponse, error) {
	lowerUsername := strings.ToLower(input.Username)
	account, err := ts.userRepo.GetUserByUsername(&model.UserList{Username: lowerUsername})
	if err != nil {
		return nil, err
	}

	passwordOk, err := ts.hasher.IsEqual(account.Pass, input.Pass)

	if err != nil {
		return nil, err
	}

	if !passwordOk {
		return nil, errs.PasswordDoesntMatch
	}

	userClaims := model.UserList{}

	if account != nil {
		userClaims.UserId = account.UserId
		userClaims.Username = account.Username
	}

	return ts.generateLoginResponse(&userClaims)

}

func (ts authService) SignUp(input dto.User) (*dto.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Pass), 10)

	if err != nil {
		return nil, err
	}
	res, err := ts.userRepo.CreateUser(&model.UserList{Username: input.UserName, Pass: string(hash)})

	if err != nil {
		return nil, err
	}
	resp := &dto.User{
		UserName: res.Username,
	}

	return resp, err
}

func (ts authService) generateLoginResponse(user *model.UserList) (*dto.LoginResponse, error) {
	accesToken, err := ts.jwtProvider.GenerateAccessToken(*user)

	if err != nil {
		return nil, err
	}

	refreshToken, err := ts.jwtProvider.GenerateRefreshToken(*user)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccesToken:   accesToken,
		RefreshToken: refreshToken,
	}, nil
}
