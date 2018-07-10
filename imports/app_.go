package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/webcore/core"

import "github.com/shaomingquan/webcore-sample/apps"

import middwares "github.com/shaomingquan/webcore-sample/middwares"

func Start_(app *core.App) {

	app.MidWare(
		"/",
		middwares.Demo(`root`),
	)

	app.Router(
		"/",
		apps.MethodOfRoot,
		apps.PrefixOfRoot,

		apps.HandlerOfRoot,
		func(ctx *gin.Context) { ctx.Next() },
	)

}

// auto generate by _, dont modify
