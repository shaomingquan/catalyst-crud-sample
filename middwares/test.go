package middwares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Demo(word string) gin.HandlerFunc {
	// initial with string param
	return func(ctx *gin.Context) {
		fmt.Println("START----" + word + "----")
		ctx.Next()
		fmt.Println("END----" + word + "----")
	}
}
