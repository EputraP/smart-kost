package service

import (
	"net/http"
	"smart-kost-backend/repository"
)

type UserCurrentLocationService interface {
	UpdateUserCurrentLocation()
}

type userCurrentLocationService struct {
	userCurrentLocationRepo repository.UserCurrentLocationRepo
}

type UserCurrentLocationServiceConfig struct {
	UserCurrentLocationRepo repository.UserCurrentLocationRepo
}

func NewUserCurrentLocationService(config UserCurrentLocationServiceConfig) UserCurrentLocationService {
	return &userCurrentLocationService{
		userCurrentLocationRepo: config.UserCurrentLocationRepo,
	}
}

func (s userCurrentLocationService) UpdateUserCurrentLocation() {
	GetAddressFromLatLong()
	// var resp *dto.UserResponse
	// res, err := s.userRepo.UpdateIsOnline(&model.UserList{UserId: input.UserId, IsOnline: input.IsOnline})

	// if err != nil {
	// 	println(err)
	// }

	// resp = &dto.UserResponse{
	// 	UserId:   res.UserId,
	// 	Username: res.Username,
	// 	IsOnline: res.IsOnline,
	// 	IsSOS:    res.IsSOS,
	// }

	// return resp, nil
}
func GetAddressFromLatLong() {
	apiUrl := "https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=-6.927770970656957&lon=107.61239014399705"
	res, err := http.Get(apiUrl)
	if res != nil {
		println(err)
	}
	println(res)
}
