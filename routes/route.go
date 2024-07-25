package routes

import (
	"smart-kost-backend/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Test       *handler.TestHandler
	HumTempRaw *handler.HumTempRawHandler
	Auth       *handler.AuthHandler
	User       *handler.UserHandler
}
type Middlewares struct {
	Auth gin.HandlerFunc
}

func Build(srv *gin.Engine, h Handlers, middlewares Middlewares) {
	auth := srv.Group("/auth")
	auth.POST("/register", h.Auth.SignUp)
	auth.POST("/login", h.Auth.Login)
	auth.GET("/logout/:userid", h.Auth.Logout)

	user := srv.Group("/user")
	user.PUT("/update-online", middlewares.Auth, h.User.UpdateOnline)
	user.PUT("/update-sos", middlewares.Auth, h.User.UpdateSOS)

	test := srv.Group("test")
	test.GET("/", middlewares.Auth, h.Test.Test)

	humTempRaw := srv.Group("/humTempRaw")
	humTempRaw.POST("/create", h.HumTempRaw.CreateHumTempRaw)

}
