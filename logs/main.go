package main

import (
	"logs/logger"
	"time"
)

func main() {
	time.Sleep(time.Second)
	logger.Logger.Info("Program started")

	for {
		logger.Logger.Debug("Running", "timestamp", time.Now().Unix())
		time.Sleep(time.Second)
	}
}
