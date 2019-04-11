package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	core "github.com/shaomingquan/catalyst/core"
)

var app = core.App{}
var prepare = make(chan int)
var loaded = make(chan int)

func init() {
	port := 0
	flag.IntVar(&port, "port", 0, "app start port")
	flag.CommandLine.Parse(os.Args[1:])

	confbyted, err := ioutil.ReadFile("./appconf.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	conf := core.Conf{}
	json.Unmarshal(confbyted, &conf)
	if port != 0 {
		conf.Port = port
	}

	app = core.App{
		Config: &conf,
	}
	prepare <- 1
}

func main() {
	<-loaded
	app.Start()
}
