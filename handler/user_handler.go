package handler

import (
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
	// var loginUser dto.LoginBody

	// if err := c.ShouldBindJSON(&loginUser); err != nil {
	// 	response.Error(c, 400, errs.InvalidRequestBody.Error())
	// 	return
	// }

	h.userService.UpdateOnline()

	// if err != nil {
	// 	if errors.Is(err, errs.PasswordDoesntMatch) ||
	// 		errors.Is(err, gorm.ErrRecordNotFound) {
	// 		response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
	// 		return
	// 	}

	// 	response.UnknownError(c, err)
	// 	return
	// }

	response.JSON(c, 201, "Update online status success", "")

}

func (h UserHandler) UpdateSOS(c *gin.Context) {
	// var loginUser dto.LoginBody

	// if err := c.ShouldBindJSON(&loginUser); err != nil {
	// 	response.Error(c, 400, errs.InvalidRequestBody.Error())
	// 	return
	// }

	h.userService.UpdateSOS()

	// if err != nil {
	// 	if errors.Is(err, errs.PasswordDoesntMatch) ||
	// 		errors.Is(err, gorm.ErrRecordNotFound) {
	// 		response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
	// 		return
	// 	}

	// 	response.UnknownError(c, err)
	// 	return
	// }

	response.JSON(c, 201, "Login success", "")

}
