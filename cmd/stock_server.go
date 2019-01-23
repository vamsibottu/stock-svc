package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/stock-svc/api/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HealthCheck)
	r.HandleFunc("/stock/{symbol}", handlers.StocksHandler)
	fmt.Println("Starting server on Port 8000")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
