package utils

import (
	log "github.com/sirupsen/logrus"
)


func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Info("Logger initialized")
}
