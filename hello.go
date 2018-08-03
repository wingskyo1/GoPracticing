package main

import (
	//"hello/app"
	"hello/app"
	"hello/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")

}
