package service

import (
	"encoding/json"
	"io"
	"net/http"
	"smart-kost-backend/dto"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
	"strconv"
	"time"
)

type UserCurrentLocationService interface {
	UpdateSOS(input dto.UpdateUserSOS) (*dto.UpdateUserSOS, error)
	UpdateUserCurrentLocation(input dto.UpdateUserLocation) (*dto.UserCurrentLocation, error)
	GetUserCurrentLocation() ([]*dto.GetUserCurrentLocationResponse, error)
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

func (s userCurrentLocationService) UpdateSOS(input dto.UpdateUserSOS) (*dto.UpdateUserSOS, error) {
	var resp *dto.UpdateUserSOS
	res, err := s.userCurrentLocationRepo.UpdateIsSOS(&model.UserCurrentLocation{UserId: input.UserId, IsSOS: input.IsSOS})

	if err != nil {
		println(err)
	}

	resp = &dto.UpdateUserSOS{
		UserId: res.UserId,
		IsSOS:  res.IsSOS,
	}

	return resp, nil
}

func (s userCurrentLocationService) GetUserCurrentLocation() ([]*dto.GetUserCurrentLocationResponse, error) {
	var resp []*dto.GetUserCurrentLocationResponse

	res, err := s.userCurrentLocationRepo.GetCurrentUserData()
	if err != nil {
		return nil, err
	}

	for _, value := range res {

		address, err := GetAddressFromLatLong(value.Lat, value.Long)

		if err != nil {
			return nil, err
		}
		longInt, err := strconv.ParseFloat(value.Long, 64)
		if err != nil {
			return nil, err
		}
		latInt, err := strconv.ParseFloat(value.Lat, 64)
		if err != nil {
			return nil, err
		}
		userId , err :=strconv.Atoi(value.UserId,)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &dto.GetUserCurrentLocationResponse{
			UserId: userId ,
			Username:   value.Username,
			IsOnline:   value.IsOnline,
			IsSOS:      value.IsSOS,
			StatusName: value.StatusName,
			Long:       longInt,
			Lat:        latInt,
			Address:    address.DisplayName,
			IconColor:  value.IconColor,
		})
		time.Sleep(2 * time.Second)
	}
	return resp, nil
}

func GetAddressFromLatLong(lat string, long string) (*dto.GetLocation, error) {

	var locationData *dto.GetLocation

	if len(lat) != 0 && len(long) != 0 {
		accessToken := "pk.6006ccd0ecc3de5a24ef402578b09b08"
		// apiUrl := "https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=" + lat + "&lon=" + long
		apiUrl := "https://us1.locationiq.com/v1/reverse?key=" + accessToken + "&lat=" + lat + "&lon=" + long + "&format=json&"
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
	} else {
		var address *dto.Address = &dto.Address{
			Road:        "",
			Subdistrict: "",
			City:        "",
			Province:    "",
		}

		locationData = &dto.GetLocation{
			DisplayName: "",
			Address:     *address,
		}
	}

	return locationData, nil
}
