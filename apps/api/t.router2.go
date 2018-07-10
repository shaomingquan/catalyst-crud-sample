package api

import (
	"github.com/gin-gonic/gin"
)

var DecoratorOfWorld = []string{
	"middwares@Demo#root-api-world-decorator", // pkg@method#param1,param2
}

func HandlerOfWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "this is world of api",
	})
}
