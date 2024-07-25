package service

import (
	"smart-kost-backend/repository"
)

type UserService interface {
	UpdateOnline()
	UpdateSOS()
}

type userService struct {
	userRepo repository.UserRepository
}

type UserServiceConfig struct {
	UserRepo repository.UserRepository
}

func NewUserService(config UserServiceConfig) UserService {
	return &userService{
		userRepo: config.UserRepo,
	}
}

func (s userService) UpdateOnline() {
	println("Update online")
}
func (s userService) UpdateSOS() {
	println("Update SOS")
}
