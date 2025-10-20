package main

import (
	"github.com/fire9900/go-forum/config"
	"github.com/fire9900/go-forum/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
