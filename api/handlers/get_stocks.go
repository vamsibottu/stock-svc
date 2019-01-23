package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stock-svc/service"
)

// NewStock creates a new application level runtime that encapsulates the shared services for this application
func NewStock() (service.Stocks, error) {
	stockservice, err := service.NewStock()
	if err != nil {
		return nil, err
	}
	return stockservice, nil
}

// StocksHandler is used to retrive stock price details of given symbol
func StocksHandler(w http.ResponseWriter, r *http.Request) {

	// capture the symbols from URL
	vars := mux.Vars(r)
	symbols := vars["symbol"]
	// capture exchange value from URL, its an optional value
	exchanges := r.FormValue("stock_exchange")

	if len(symbols) < 3 {
		http.Error(w, fmt.Sprintln("invalid symbol"), http.StatusBadRequest)
	}

	stocksvc, err := NewStock()
	if err != nil {
		fmt.Fprintln(w, "error initiating stock service")
	}

	stockdetails, err := stocksvc.GetStocks(symbols, exchanges)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp, err := json.Marshal(stockdetails)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	fmt.Fprintln(w, string(resp))
}
