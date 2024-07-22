package routes

import (
	"smart-kost-backend/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Test       *handler.TestHandler
	HumTempRaw *handler.HumTempRawHandler
	User       *handler.UserHandler
}
type Middlewares struct {
	Auth gin.HandlerFunc
}

func Build(srv *gin.Engine, h Handlers, middlewares Middlewares) {
	auth := srv.Group("/auth")
	auth.POST("/register", h.User.SignUp)
	auth.POST("/login", h.User.Login)

	test := srv.Group("test")
	test.GET("/", middlewares.Auth, h.Test.Test)

	humTempRaw := srv.Group("/humTempRaw")
	humTempRaw.POST("/create", h.HumTempRaw.CreateHumTempRaw)

}
