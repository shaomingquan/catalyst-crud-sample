package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type connection struct {
	user     string
	password string
	host     string
	dbname   string
}

func (c *connection) makeConnectionStr() string {
	return c.user +
		":" + c.password +
		"@tcp(" + c.host +
		")/" + c.dbname +
		"?charset=utf8&parseTime=True&loc=Local"
}

func (c *connection) getDBObj(t string) *gorm.DB {
	db, err := gorm.Open(t, c.makeConnectionStr())
	if err != nil {
		panic(err) // how to
	}
	db = db.Debug()
	return db
}

var localhostdb = connection{
	"root",
	"123456",
	"127.0.0.1:3333",
	"test",
}

var (
	LocalhostDB *gorm.DB
)

func init() {
	LocalhostDB = localhostdb.getDBObj("mysql")
}

func GetDB() *gorm.DB {
	return LocalhostDB
}
