package handler

import (
	"smart-kost-backend/service"

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

	// resp, err := h.Service.Login(loginUser)

	// if err != nil {
	// 	if errors.Is(err, errs.PasswordDoesntMatch) ||
	// 		errors.Is(err, gorm.ErrRecordNotFound) {
	// 		response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
	// 		return
	// 	}

	// 	response.UnknownError(c, err)
	// 	return
	// }

	// response.JSON(c, 200, "Login success", resp)

}

func (h UserHandler) UpdateSOS(c *gin.Context) {
	// var loginUser dto.LoginBody

	// if err := c.ShouldBindJSON(&loginUser); err != nil {
	// 	response.Error(c, 400, errs.InvalidRequestBody.Error())
	// 	return
	// }

	// resp, err := h.Service.Login(loginUser)

	// if err != nil {
	// 	if errors.Is(err, errs.PasswordDoesntMatch) ||
	// 		errors.Is(err, gorm.ErrRecordNotFound) {
	// 		response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
	// 		return
	// 	}

	// 	response.UnknownError(c, err)
	// 	return
	// }

	// response.JSON(c, 200, "Login success", resp)

}
