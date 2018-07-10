package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/webcore/core"

import "net/http"
import validator "gopkg.in/validator.v2"
import gene "github.com/shaomingquan/webcore/gene"

import "github.com/shaomingquan/webcore-curd-sample/apps/api/article"

import middwares "github.com/shaomingquan/webcore-curd-sample/middwares"

func Start_api_article(app *core.App) {

	var paramsValidatorOfGetDetail = func(
		ctx *gin.Context,
	) {
		// validate
		paramsInstance := article.ParamsOfGetDetail{
			gene.ParamTostring(ctx, "article_id"),
			gene.ParamTofloat64(ctx, "version"),
		}
		if err := validator.Validate(paramsInstance); err != nil {
			ctx.JSON(400, err.Error())
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			ctx.Set("xParams", &paramsInstance)
			ctx.Next()
		}
	}

	app.MidWare(
		"/api/article",
		middwares.Demo(`root-api-article`),
	)

	app.Router(
		"/api/article",
		article.MethodOfGetDetail,
		article.PrefixOfGetDetail,

		paramsValidatorOfGetDetail,
		article.HandlerOfGetDetail,
	)

}

// auto generate by _, dont modify
