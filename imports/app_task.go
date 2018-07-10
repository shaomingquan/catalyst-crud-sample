package imports

import "github.com/gin-gonic/gin"
import core "github.com/shaomingquan/webcore/core"

import "github.com/shaomingquan/webcore-sample/apps/task"

import middwares "github.com/shaomingquan/webcore-sample/middwares"

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
