package main

import (
	"flag"
	"log"

	"github.com/Killayt/image-generator/configs"
	"github.com/Killayt/image-generator/internal/server"
)

var confPath = flag.String("conf-path", "configs/.env", "Path to config file (.env).")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatal("!!! Server is down !!!\n", err)
	}

	server.Run(conf)
}
