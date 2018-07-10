# webcore-curd-sample

> it's an extention template base on [webcore-sample](https://github.com/shaomingquan/webcore-sample)

***benifit: build basic curd api quickly***

0, prepare

install project vendors 

install mysql and create test table (./infrastructure.dev could help)

import `curd demo.postman_collection.json` with postman, see request samples

1, define a model and define its lifecycle:

- SingleInstance: single instance just for unmarshal
- ListInstance: return collection of test instance, just for return values
- NewInstance: when new item, add some default value
- ReturnInstance: when get list
- Put???

```go
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
			return &Test{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		},
		ReturnInstance: func(ctx *gin.Context, ret interface{}) interface{} {

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
```

2, declare curd interseptor

```go
package apps

var MiddlewaresComposer = []string{

	// curd interseptor
	"store@Curd#/api/data/test/,test",

	"middwares@Demo#root", // pkg@method#param1,param2
}

var PrefixOfRoot = "/"
var MethodOfRoot = "GET"
```