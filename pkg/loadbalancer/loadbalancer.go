package loadbalancer

import (
	"log"
	"net/http"

	"github.com/koddr/go-lb/pkg/healthcheck"
	"github.com/koddr/go-lb/pkg/serverpool"
)

// LoadBalancer load balances the incoming request
func LoadBalancer(w http.ResponseWriter, r *http.Request) {
	//
	serverpool := serverpool.NewServerPool()

	attempts := healthcheck.GetAttemptsFromContext(r)
	if attempts > 3 {
		log.Printf("%s(%s) Max attempts reached, terminating\n", r.RemoteAddr, r.URL.Path)
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		return
	}

	peer := serverpool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
