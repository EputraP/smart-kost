package service

import (
	"encoding/json"
	"io"
	"net/http"
	"smart-kost-backend/dto"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
)

type UserCurrentLocationService interface {
	UpdateUserCurrentLocation(input dto.UpdateUserLocation) (*dto.UserCurrentLocation, error)
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

func (s userCurrentLocationService) UpdateUserCurrentLocation(input dto.UpdateUserLocation) (*dto.UserCurrentLocation, error) {
	// res := GetAddressFromLatLong("-6.927770970656957", "107.61239014399705")
	// fmt.Println(res.DisplayName)
	// fmt.Println(res.Address.Road)

	var resp *dto.UserCurrentLocation
	res, err := s.userCurrentLocationRepo.UpdateLatLong(&model.UserCurrentLocation{UserId: input.UserId, UserCurrentLocationLat: input.UserCurrentLocationLat, UserCurrentLocationLong: input.UserCurrentLocationLong})

	if err != nil {
		return nil, err
	}

	resp = &dto.UserCurrentLocation{
		UserId:                  res.UserId,
		StatusId:                res.StatusId,
		UserCurrentLocationLat:  res.UserCurrentLocationLat,
		UserCurrentLocationLong: res.UserCurrentLocationLong,
	}

	return resp, nil
}

func GetAddressFromLatLong(lat string, long string) (*dto.GetLocation, error) {

	var locationData *dto.GetLocation
	apiUrl := "https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=" + lat + "&lon=" + long
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return nil, err
	}

	return locationData, nil

}
