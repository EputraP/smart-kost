package handler

import (
	"net/http"
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

func (h UserHandler) Login(c *gin.Context) {
	var loginUser dto.User
	
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.userService.Login(loginUser);

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", resp, 3600*24,"","",false,false)

	response.JSON(c, 200, "Login Success", resp)

}

func (h UserHandler) SignUp (c *gin.Context) {
	var createUser dto.User
	
	if err := c.ShouldBindJSON(&createUser); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.userService.SignUp(createUser);

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create User Success", resp)
}


