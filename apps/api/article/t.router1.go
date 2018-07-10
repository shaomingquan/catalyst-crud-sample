package article

import (
	"github.com/gin-gonic/gin"
)

// http://localhost:7777/api/article/detail?version=1.2&article_id=abcde

type ParamsOfGetDetail struct {
	ArticleId string  `validate:"nonzero"`
	Version   float64 `validate:"nonzero"`
}

func HandlerOfGetDetail(ctx *gin.Context) {
	_params, _ := ctx.Get("xParams")
	params := _params.(*ParamsOfGetDetail)
	ctx.JSON(200, gin.H{
		"message": "text book",
		"xParams": params,
	})
}
