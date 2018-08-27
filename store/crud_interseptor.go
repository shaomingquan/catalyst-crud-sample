package store

import (
	"encoding/json"
	"net/http"
	"strings"

	errs "errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Crud(prefix, model string) gin.HandlerFunc {

	instanceGenerator := modelInstanceMapper[model]

	return func(ctx *gin.Context) {

		db := GetDB()

		println(db == nil)

		path := ctx.Request.URL.Path
		println(prefix)
		println(ctx.Request.URL.Path)
		println(strings.HasPrefix(path, prefix))
		if strings.HasPrefix(path, prefix) {
			method := ctx.Request.Method
			switch method {
			case "GET":
				single := instanceGenerator.SingleInstance(ctx)
				list := instanceGenerator.ListInstance(ctx)
				err := DoGET(ctx, db, single, list)

				if err != nil {
					HttpEndWith500(ctx, errs.New("db error"))
					return
				}

				single2 := instanceGenerator.SingleInstance(ctx)
				list2 := instanceGenerator.ListInstance(ctx)
				count, page, err := DoCount(ctx, db, single2, list2)

				if err != nil {
					HttpEndWith500(ctx, errs.New("db error"))
					return
				}

				ctx.JSON(200, gin.H{
					"data":  instanceGenerator.ReturnInstance(ctx, list),
					"count": count,
					"page":  page,
				})
			case "POST":
				newinstance := instanceGenerator.NewInstance(ctx)
				err := DoPOST(ctx, db, newinstance)

				if err != nil {
					HttpEndWith400(ctx, err)
					return
				}

				ctx.JSON(200, gin.H{
					"newinstance": newinstance,
				})

			case "PUT":
				single := instanceGenerator.SingleInstance(ctx)

				err := DoPUT(ctx, db, single)

				if err != nil {
					HttpEndWith400(ctx, err)
					return
				}

				ctx.JSON(200, gin.H{
					"update": single,
				})

			case "DELETE":
				single := instanceGenerator.SingleInstance(ctx)

				err := DoDELETE(ctx, db, single)

				if err != nil {
					HttpEndWith400(ctx, err)
					return
				}

				ctx.JSON(200, gin.H{
					"deletedinstance": single,
				})
			}

		} else {
			ctx.Next()
		}

	}
}

func DoCount(ctx *gin.Context, db *gorm.DB, single, list interface{}) (int, int, error) {
	optionsStr := ctx.Query("options")
	queryStr := ctx.Query("query")

	json.Unmarshal([]byte(queryStr), single) // get

	options := GetOptions{}
	json.Unmarshal([]byte(optionsStr), &options)

	page := options.Page

	count, err := Count(
		db,
		single,
		&options,
		list,
	)

	return count, page, err
}

func DoGET(ctx *gin.Context, db *gorm.DB, single, list interface{}) error {
	optionsStr := ctx.Query("options")
	queryStr := ctx.Query("query")

	json.Unmarshal([]byte(queryStr), single) // get

	options := GetOptions{}
	json.Unmarshal([]byte(optionsStr), &options)

	fieldsStr := ctx.Query("fields")

	fields := []string{}
	if fieldsStr != "" {
		json.Unmarshal([]byte(fieldsStr), &fields)
	}

	err := Get(
		fields,
		db,
		single, // query
		&options,
		list,
	)
	return err

}

func DoPOST(ctx *gin.Context, db *gorm.DB, newinstance interface{}) error {

	jsonBody, _ := ctx.GetRawData() // json in post body

	json.Unmarshal(jsonBody, newinstance)

	err := Post(db, newinstance)

	return err
}

func DoPUT(ctx *gin.Context, db *gorm.DB, query interface{}) error {
	// insert if first
	queryStr := ctx.Query("query")

	hasid := checkId([]byte(queryStr))

	if !hasid {
		return errs.New("no id")
	}

	json.Unmarshal([]byte(queryStr), query) // get

	newinstance := map[string]interface{}{}
	jsonBody, _ := ctx.GetRawData() // json in post body
	json.Unmarshal(jsonBody, &newinstance)

	err := Put(db, query, newinstance)

	return err
}

func DoDELETE(ctx *gin.Context, db *gorm.DB, query interface{}) error {
	jsonBody, _ := ctx.GetRawData() // json in post body

	hasid := checkId(jsonBody)

	if !hasid {
		return errs.New("no id")
	}

	json.Unmarshal(jsonBody, query)

	err := Delete(db, query)

	return err
}

func checkId(body []byte) bool {
	ret := map[string]interface{}{}
	json.Unmarshal(body, &ret)

	id, ok := ret["id"]
	if !ok {
		return false
	}

	_, ok = id.(string)
	if ok {
		return false
	}

	_, ok = id.(float64)

	return ok
}

func HttpEndWith500(ctx *gin.Context, err error) {
	ctx.JSON(500, gin.H{
		"message": err.Error(),
	})
	ctx.AbortWithStatus(http.StatusInternalServerError)
}

func HttpEndWith400(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"message": err.Error(),
	})
	ctx.AbortWithStatus(http.StatusBadRequest)
}
