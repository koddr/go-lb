package healthcheck

import (
	"log"
	"net/http"
	"time"

	"github.com/koddr/go-lb/pkg/serverpool"
)

const (
	Attempts int = iota
	Retry
)

// GetAttemptsFromContext returns the attempts for request
func GetAttemptsFromContext(r *http.Request) int {
	if attempts, ok := r.Context().Value(Attempts).(int); ok {
		return attempts
	}
	return 1
}

// GetAttemptsFromContext returns the attempts for request
func GetRetryFromContext(r *http.Request) int {
	if retry, ok := r.Context().Value(Retry).(int); ok {
		return retry
	}
	return 0
}

// Start runs a routine for check status of the backends every 2 mins
func Start() {
	//
	t := time.NewTicker(time.Minute * 2)

	//
	serverpool := serverpool.NewServerPool()

	//
	for {
		select {
		case <-t.C:
			log.Println("Starting health check...")
			serverpool.HealthCheck()
			log.Println("Health check completed")
		}
	}
}
