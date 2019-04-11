package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/catalyst/core"

import "github.com/shaomingquan/catalyst-crud-sample/apps/api"

import middwares "github.com/shaomingquan/catalyst-crud-sample/middwares"

func Start_api(app *core.App) {

	app.MidWare(
		"/api",
		middwares.Demo(`root-api`),
	)

	app.Router(
		"/api",
		api.MethodOfHello,
		api.PrefixOfHello,

		api.HandlerOfHello,
		func(ctx *gin.Context) { ctx.Next() },
	)

	app.Router(
		"/api",
		api.MethodOfWorld,
		api.PrefixOfWorld,
		middwares.Demo(`root-api-world-decorator`),

		api.HandlerOfWorld,
		func(ctx *gin.Context) { ctx.Next() },
	)

}

// auto generate by _, dont modify
