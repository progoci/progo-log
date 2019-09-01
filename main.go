package main

import (
	"path/filepath"

	"progo/core/config"
	"progo/loom/cmd/service"
)

func main() {
	envPath, _ := filepath.Abs("./.env")
	config.Init(envPath)

	port := ":" + config.Get("HOST_PORT")

	service.Run(port)
}
