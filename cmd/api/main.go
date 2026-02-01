package main

import (
	"kasir-api/internal/bootstrap"
)

func main() {
	// initialize bootstrap
	app := bootstrap.InitApp()

	// run server
	app.Run()
}
