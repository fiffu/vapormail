package main

import (
	"github.com/fiffu/vapormail/config"
	"github.com/fiffu/vapormail/services"
	"github.com/fiffu/vapormail/services/hub"
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
