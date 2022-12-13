package main

import (
	"log"

	"github.com/Killayt/image-generator/configs"
	"github.com/Killayt/image-generator/internal/server"
)

func main() {
	conf, err := configs.New("./configs/.env")
	if err != nil {
		log.Fatal("!!! Server is down !!!\n", err)
	}

	server.Run(conf)
}
