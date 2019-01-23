package handlers

import "net/http"

// HealthCheck is used to check the status of the service
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Stock Service is UP"))
}
