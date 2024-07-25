package service

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
)

type UserService interface {
	UpdateOnline(input dto.UpdateUserOnlineSOS) (*dto.UserResponse, error)
	UpdateSOS(input dto.UpdateUserOnlineSOS) (*dto.UserResponse, error)
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

func (s userService) UpdateOnline(input dto.UpdateUserOnlineSOS) (*dto.UserResponse, error) {

	var resp *dto.UserResponse
	res, err := s.userRepo.UpdateIsOnline(&model.UserList{UserId: input.UserId, IsOnline: input.IsOnline})

	if err != nil {
		println(err)
	}

	resp = &dto.UserResponse{
		UserId:   res.UserId,
		Username: res.Username,
		IsOnline: res.IsOnline,
		IsSOS:    res.IsSOS,
	}

	return resp, nil
}
func (s userService) UpdateSOS(input dto.UpdateUserOnlineSOS) (*dto.UserResponse, error) {
	var resp *dto.UserResponse
	res, err := s.userRepo.UpdateIsSOS(&model.UserList{UserId: input.UserId, IsSOS: input.IsSOS})

	if err != nil {
		println(err)
	}

	resp = &dto.UserResponse{
		UserId:   res.UserId,
		Username: res.Username,
		IsOnline: res.IsOnline,
		IsSOS:    res.IsSOS,
	}

	return resp, nil
}
