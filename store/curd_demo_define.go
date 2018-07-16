package store

import (
	"time"

	"github.com/gin-gonic/gin"
)

type InstanceLifecycle struct {
	SingleInstance func(ctx *gin.Context) interface{}
	ListInstance   func(ctx *gin.Context) interface{}

	NewInstance    func(ctx *gin.Context) interface{}
	ReturnInstance func(ctx *gin.Context, ret interface{}) interface{}
}

var modelInstanceMapper = map[string]InstanceLifecycle{
	"test": InstanceLifecycle{
		SingleInstance: func(ctx *gin.Context) interface{} {
			return &Test{}
		},
		ListInstance: func(ctx *gin.Context) interface{} {
			return &[]Test{}
		},
		NewInstance: func(ctx *gin.Context) interface{} {
			// post hook
			// maybe mysql CURRENT_TIME, just for demo
			return &Test{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		},
		ReturnInstance: func(
			ctx *gin.Context,
			ret interface{},
		) interface{} {
			// get hook
			// formatter, parser, mixin other data...
			_list := *ret.(*[]Test)
			list := make([]Test, len(_list))

			for index, item := range _list {
				item.Name = item.Name + " ~~"
				list[index] = item
			}
			return list
		},
	},
}
