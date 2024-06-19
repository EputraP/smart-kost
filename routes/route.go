package routes

import (
	"smart-kost-backend/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Test       *handler.TestHandler
	HumTempRaw *handler.HumTempRawHandler
}

func Build(srv *gin.Engine, h Handlers) {
	test := srv.Group("test")
	test.GET("/", h.Test.Test)

	humTempRaw := srv.Group("/humTempRaw")
	humTempRaw.POST("/create", h.HumTempRaw.CreateHumTempRaw)

}
