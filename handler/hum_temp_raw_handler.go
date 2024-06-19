package handler

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"

	"github.com/gin-gonic/gin"
)

type HumTempRawHandler struct {
	humTempRawService service.HumTempRawService
}

type HumTempRawHandlerConfig struct {
	HumTempRawService service.HumTempRawService
}

func NewHumTempRawHandler(config HumTempRawHandlerConfig) *HumTempRawHandler {
	return &HumTempRawHandler{
		humTempRawService: config.HumTempRawService,
	}
}

func (h HumTempRawHandler) CreateHumTempRaw(c *gin.Context) {
	var createHumTempRaw dto.CreateHumTempRaw

	if err := c.ShouldBindJSON(&createHumTempRaw); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.humTempRawService.CreateHumTempRaw(createHumTempRaw)

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create HumTempRaw Success", resp)
}
