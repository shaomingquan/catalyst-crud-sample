package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/catalyst/core"

import "github.com/shaomingquan/catalyst-crud-sample/apps/task"

import middwares "github.com/shaomingquan/catalyst-crud-sample/middwares"

func Start_task(app *core.App) {

	app.MidWare(
		"/task",
		middwares.Demo(`root-task`),
	)

	app.Router(
		"/task",
		task.MethodOfHello,
		task.PrefixOfHello,

		task.HandlerOfHello,
		func(ctx *gin.Context) { ctx.Next() },
	)

	app.Router(
		"/task",
		task.MethodOfWorld,
		task.PrefixOfWorld,

		task.HandlerOfWorld,
		func(ctx *gin.Context) { ctx.Next() },
	)

}

// auto generate by _, dont modify
