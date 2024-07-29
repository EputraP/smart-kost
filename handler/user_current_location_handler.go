package handler

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCurrentLocationHandler struct {
	userCurrentLocationService service.UserCurrentLocationService
}

type UserCurrentLocationHandlerConfig struct {
	UserCurrentLocationService service.UserCurrentLocationService
}

func NewUserCurrentLocationRawHandler(config UserCurrentLocationHandlerConfig) *UserCurrentLocationHandler {
	return &UserCurrentLocationHandler{
		userCurrentLocationService: config.UserCurrentLocationService,
	}
}

func (h UserCurrentLocationHandler) UpdateSOS(c *gin.Context) {
	var userSOS dto.UpdateUserSOS

	if err := c.ShouldBindJSON(&userSOS); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.userCurrentLocationService.UpdateSOS(userSOS)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Update SOS status success", resp)

}

func (h UserCurrentLocationHandler) UpdateUserCurrentLocation(c *gin.Context) {
	var userLocation dto.UpdateUserLocation

	if err := c.ShouldBindJSON(&userLocation); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}
	resp, err := h.userCurrentLocationService.UpdateUserCurrentLocation(userLocation)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 201, "Update online status success", resp)

}
func (h UserCurrentLocationHandler) GetCurrentUserLocation(c *gin.Context) {
	resp, err := h.userCurrentLocationService.GetUserCurrentLocation()
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Get current user data success", resp)

}

func (h UserCurrentLocationHandler) GetSingleCurrentUserLocation(c *gin.Context) {
	if len(c.Param("userid")) == 0 {
		response.Error(c, 400, errs.InvalidRequestParam.Error())
		return
	}
	idStr, _ := strconv.Atoi(c.Param("userid"))
	resp, err := h.userCurrentLocationService.GetUserCurrentLocationByUserId(idStr)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Get single current user data success", resp)

}
