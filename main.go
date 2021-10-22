package main

import (
	"github.com/fiffu/vprmail/config"
	"github.com/fiffu/vprmail/services"
	"github.com/fiffu/vprmail/services/hub"
	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	log.SetReportCaller(true)
}

func main() {
	setupLogger()
	cfg := config.SetupConfig()
	log.Infof("Config=%+v", cfg)

	hub := hub.NewHubService(cfg)

	// HTTP API server
	httpAPI := services.NewHTTPService(cfg.APIPort, hub)
	httpAPI.Start()
}
