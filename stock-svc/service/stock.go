package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/stock-svc/models"
)

var (
	host     = "https://www.worldtradingdata.com/api/v1"
	apiToken = "2qNAyqIXfNxcMayuc5m0VqpShrgadMchDUCpQgjeU3MNprg91yXeqWjYSYbr"
)

// stockResponse is used to unmarshal the response from trading URL
type stockResponse struct {
	SymbolsRequested int    `json:"symbols_requested"`
	SymbolsReturned  int    `json:"symbols_returned"`
	Data             []data `json:"data"`
}

type data struct {
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	Currency           string `json:"currency"`
	Price              string `json:"price"`
	PriceOpen          string `json:"price_open"`
	DayHigh            string `json:"day_high"`
	DayLow             string `json:"day_low"`
	Five2WeekHigh      string `json:"52_week_high"`
	Five2WeekLow       string `json:"52_week_low"`
	DayChange          string `json:"day_change"`
	ChangePct          string `json:"change_pct"`
	CloseYesterday     string `json:"close_yesterday"`
	MarketCap          string `json:"market_cap"`
	Volume             string `json:"volume"`
	VolumeAvg          string `json:"volume_avg"`
	Shares             string `json:"shares"`
	StockExchangeLong  string `json:"stock_exchange_long"`
	StockExchangeShort string `json:"stock_exchange_short"`
	Timezone           string `json:"timezone"`
	TimezoneName       string `json:"timezone_name"`
	GmtOffset          string `json:"gmt_offset"`
	LastTradeTime      string `json:"last_trade_time"`
}

// GetStocks is the method used to retrieve stock prices of a given symbol
func (s *stocks) GetStocks(symbols, exchanges string) (map[string][]models.StockDetails, error) {

	// capture response in stocks variable
	stocks := map[string][]models.StockDetails{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/stock?symbol=%s&api_token=%s", host, symbols, apiToken), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to retrive stock details")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := &stockResponse{}
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	if resp != nil && len(resp.Data) > 0 {
		for _, sd := range resp.Data {
			var ex string
			if strings.Contains(exchanges, sd.StockExchangeShort) {
				ex = sd.StockExchangeShort
			}
			stockDeatils := models.StockDetails{
				Symbol:         sd.Symbol,
				Name:           sd.Name,
				Price:          sd.Price,
				CloseYesterday: sd.CloseYesterday,
				MarketCap:      sd.MarketCap,
				Volume:         sd.Volume,
				TimeZone:       sd.Timezone,
				TimeZoneName:   sd.TimezoneName,
				GmtOffset:      sd.GmtOffset,
				LastTradeTime:  sd.LastTradeTime,
			}
			stocks[ex] = append(stocks[ex], stockDeatils)
		}
	}
	return stocks, nil
}
