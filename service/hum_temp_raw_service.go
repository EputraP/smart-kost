package service

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/model"
	"smart-kost-backend/repository"
)

type HumTempRawService interface {
	CreateHumTempRaw(input dto.CreateHumTempRaw) (*dto.CreateHumTempRaw, error)
}

type humTempRawService struct {
	humTempRawRepo repository.HumTempRawRepository
}

type HumTempRawServiceConfig struct {
	HumTempRawRepo repository.HumTempRawRepository
}

func NewHumTempRawService(config HumTempRawServiceConfig) HumTempRawService {
	return &humTempRawService{
		humTempRawRepo: config.HumTempRawRepo,
	}
}

func (s *humTempRawService) CreateHumTempRaw(input dto.CreateHumTempRaw) (*dto.CreateHumTempRaw, error) {
	var err error
	var resp *dto.CreateHumTempRaw
	var res *model.HumTemRaw

	res, err = s.humTempRawRepo.CreateHumTempRawData(&model.HumTemRaw{SensorID: input.SensorID, Hum: input.Hum, Tem: input.Tem})

	if err != nil {
		return nil, err
	}
	resp = &dto.CreateHumTempRaw{
		SensorID: res.SensorID,
		Hum:      res.Hum,
		Tem:      res.Tem,
	}

	return resp, nil
}
