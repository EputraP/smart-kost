package handler

import (
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testService service.TestService
}

type TestHandlerConfig struct {
	TestService service.TestService
}

func NewTestHandler(config TestHandlerConfig) *TestHandler {
	return &TestHandler{
		testService: config.TestService,
	}
}

func (h TestHandler) Test(c *gin.Context) {
	resp := h.testService.Test()
	response.JSON(c, 200, "test success", resp)
}
