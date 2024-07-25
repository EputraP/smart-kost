package service

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/repository"
	"smart-kost-backend/util/hasher"
	"smart-kost-backend/util/tokenprovider"
)

type UserService interface {
	Login(input dto.LoginBody) (*dto.LoginResponse, error)
	SignUp(input dto.User) (*dto.User, error)
}

type userService struct {
	userRepo    repository.UserRepository
	hasher      hasher.Hasher
	jwtProvider tokenprovider.JWTTokenProvider
}

type UserServiceConfig struct {
	UserRepo    repository.UserRepository
	Hasher      hasher.Hasher
	JwtProvider tokenprovider.JWTTokenProvider
}

func NewUserService(config UserServiceConfig) UserService {
	return &authService{
		userRepo:    config.UserRepo,
		hasher:      config.Hasher,
		jwtProvider: config.JwtProvider,
	}
}

func (s userService) UpdateOnline() {
	println("Update online")
}
func (s userService) UpdateSOS() {
	println("Update SOS")
}
