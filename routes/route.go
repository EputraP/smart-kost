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

func Build(srv *gin.Engine, h Handlers) {
	test := srv.Group("test")
	test.GET("/", h.Test.Test)

	humTempRaw := srv.Group("/humTempRaw")
	humTempRaw.POST("/create", h.HumTempRaw.CreateHumTempRaw)

	user := srv.Group("/user")
	user.POST("/login", h.User.Login)
	user.POST("/signup", h.User.SignUp)
}
