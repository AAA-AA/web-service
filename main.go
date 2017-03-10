package main

import (
	"flag"
	"log"

	"web-service/conf"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Println("conf has been inited successfully!")
}
