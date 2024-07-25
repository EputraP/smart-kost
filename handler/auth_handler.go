package handler

import (
	"errors"
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService service.AuthService
}

type AuthHandlerConfig struct {
	AuthService service.AuthService
}

func NewAuthRawHandler(config AuthHandlerConfig) *AuthHandler {
	return &AuthHandler{
		authService: config.AuthService,
	}
}

func (h AuthHandler) Login(c *gin.Context) {
	var loginUser dto.LoginBody

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.authService.Login(loginUser)

	if err != nil {
		if errors.Is(err, errs.PasswordDoesntMatch) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
			return
		}

		response.UnknownError(c, err)
		return
	}

	response.JSON(c, 200, "Login success", resp)

}

func (h AuthHandler) SignUp(c *gin.Context) {
	var createUser dto.User

	if err := c.ShouldBindJSON(&createUser); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.authService.SignUp(createUser)

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create User Success", resp)
}

func (h AuthHandler) Logout(c *gin.Context) {
	if len(c.Param("userid")) == 0 {
		response.Error(c, 400, errs.InvalidRequestParam.Error())
		return
	}

	idStr, _ := strconv.Atoi(c.Param("userid"))

	resp, err := h.authService.Logout(idStr)

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Logout success", resp)
}
