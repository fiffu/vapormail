package main

import (
	"github.com/fiffu/vprmail/config"
	"github.com/fiffu/vprmail/services"
	"github.com/fiffu/vprmail/services/hub"
)

func main() {
	cfg := config.SetupConfig()

	hub := hub.NewHubService(cfg)

	// HTTP API server
	httpAPI := services.NewHTTPService(cfg.APIPort, hub)
	httpAPI.Start()
}
