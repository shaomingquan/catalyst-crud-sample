package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/catalyst/core"

import "github.com/shaomingquan/catalyst-crud-sample/apps"

import store "github.com/shaomingquan/catalyst-crud-sample/store"

import middwares "github.com/shaomingquan/catalyst-crud-sample/middwares"

func Start_(app *core.App) {

	app.MidWare(
		"/",
		store.Crud(`/api/data/test/`, `test`),
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
