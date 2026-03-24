package main

import (
	"workshop-storage-api-docs/config"
	"workshop-storage-api-docs/internal/app"
)

func main() {
	config.NewConfig()

	app.Run()
}
