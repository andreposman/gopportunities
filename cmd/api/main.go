package main

import (
	"github.com/andreposman/gopportunities/cmd/api/router"
	"github.com/andreposman/gopportunities/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.InitLogger("- main package - ")

	logger.Info("Gopportunities API")

	if err := config.Init(); err != nil {
		logger.Errf("Config initialization err: %v", err)
		panic(err)
	}

	router.Init()

}
