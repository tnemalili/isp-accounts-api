package main

import (
	log "github.com/sirupsen/logrus"
	"isp.accounts.api/routes"
)

func main() {
	// SETTING LOGGER
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, ForceColors: true})
	err := routes.ConfigureRouter()
	if err != nil { log.Errorf("[ISP.ACCOUNTS.API.main] Error: %v", err.Error())}
}