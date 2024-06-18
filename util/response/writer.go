package response

import "github.com/gin-gonic/gin"

func JSON(ctx *gin.Context, statusCode int, message string, data interface{}) {
	resp := NewResponse(statusCode, message, data)

	ctx.JSON(statusCode, resp)

}
