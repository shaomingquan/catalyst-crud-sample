package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/webcore/core"

import "github.com/shaomingquan/webcore-curd-sample/apps"

import store "github.com/shaomingquan/webcore-curd-sample/store"

import middwares "github.com/shaomingquan/webcore-curd-sample/middwares"

func Start_(app *core.App) {

	app.MidWare(
		"/",
		store.Curd(`/api/data/test/`, `test`),
	)

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
