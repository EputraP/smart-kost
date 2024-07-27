package handler

import (
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

type UserHandlerConfig struct {
	UserService service.UserService
}

func NewUserRawHandler(config UserHandlerConfig) *UserHandler {
	return &UserHandler{
		userService: config.UserService,
	}
}

func (h UserHandler) UpdateOnline(c *gin.Context) {
	var userOnline dto.UpdateUserOnline

	if err := c.ShouldBindJSON(&userOnline); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.userService.UpdateOnline(userOnline)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Update online status success", resp)

}
