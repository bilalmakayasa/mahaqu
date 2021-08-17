package main

import (
	"mahaqu/src/config"
	"mahaqu/src/routes"
)

func main() {
	config.ConnetDB()
	e := routes.RegisterRouter()
	e.Run(`:8000`)
}
