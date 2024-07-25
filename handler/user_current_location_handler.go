package handler

import (
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"

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

func (h UserCurrentLocationHandler) UpdateUserCurrentLocation(c *gin.Context) {

	h.userCurrentLocationService.UpdateUserCurrentLocation()

	response.JSON(c, 201, "Update online status success", "")

}
