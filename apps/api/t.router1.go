package api

import (
	"github.com/gin-gonic/gin"
)

func HandlerOfHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "this is hello of api",
	})
}
