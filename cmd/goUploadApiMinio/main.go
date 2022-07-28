// This file is generated using ucnbrew tool.
// Check out for more info "https://github.com/saucon/ucnbrew"
package main

import (
	"github.com/saucon/go-upload-api-minio/configs/env"
	"github.com/saucon/go-upload-api-minio/router"

	"log"
)

func main() {
	env.NewEnv(".env")
	cfg := env.Config

	router := router.NewRouter()
	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatal("Error running router : ", err)
	}
}
