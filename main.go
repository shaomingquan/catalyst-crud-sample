package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	core "github.com/shaomingquan/webcore/core"
)

var app = core.App{}
var prepare = make(chan int)
var loaded = make(chan int)

func init() {
	confbyted, err := ioutil.ReadFile("./appconf.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	conf := core.Conf{}
	json.Unmarshal(confbyted, &conf)
	app = core.App{
		Config: &conf,
	}
	prepare <- 1
}

func main() {
	<-loaded
	app.Start()
}
