package task

import (
	"github.com/gin-gonic/gin"
)

func HandlerOfWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "this is world of task",
	})
}
