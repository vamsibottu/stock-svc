package service

import (
	"net"
	"net/http"
	"time"

	"github.com/stock-svc/models"
)

// Stocks defines the available functions for the given service implementation.
type Stocks interface {
	// GetStocks is the method used to retrieve stock prices of a given symbol
	GetStocks(symbols, exchanges string) (stocks map[string][]models.StockDetails, err error)
}

type stocks struct {
	client *http.Client
}

// NewStock creates a new configured stock instance
func NewStock() (Stocks, error) {

	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	s := &stocks{
		client: client,
	}
	return s, nil
}
