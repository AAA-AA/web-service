package main

import (
	"flag"

	"web-service/conf"
	"web-service/http"

	"github.com/ngaut/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Info("conf has been inited successfully!")
	http.Init(conf.Conf)
}
