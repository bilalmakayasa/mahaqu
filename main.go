package main

import (
	"mahaqu/src/config"
	"mahaqu/src/helper"
	"mahaqu/src/routes"
)

func main() {
	helper.RedisInitialize()
	config.ConnetDB()
	e := routes.RegisterRouter()
	e.Run(`:8000`)
}
