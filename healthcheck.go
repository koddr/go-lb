package main

import (
	"log"
	"time"
)

// Start runs a routine for check status of the backends every 2 mins
func Start() {
	//
	t := time.NewTicker(time.Minute * 2)

	//
	for {
		select {
		case <-t.C:
			log.Println("Starting health check...")
			serverPool.HealthCheck()
			log.Println("Health check completed")
		}
	}
}
