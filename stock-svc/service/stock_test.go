package service

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stock-svc/models"
)

var successresp = map[string][]models.StockDetails{
	"NASDAQ": {
		{
			Symbol:         "AAPL",
			Name:           "Apple Inc.",
			Price:          "153.30",
			CloseYesterday: "156.82",
			MarketCap:      "725078814334",
			Volume:         "30393970",
			TimeZone:       "EST",
			TimeZoneName:   "America/New_York",
			GmtOffset:      "-18000",
			LastTradeTime:  "2019-01-22 16:00:01",
		},
	},
}

func Test_stocks_Service(t *testing.T) {

	positiveserv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{
			"symbols_requested": 3,
			"symbols_returned": 3,
			"data": [
				{
					"symbol": "AAPL",
					"name": "Apple Inc.",
					"currency": "USD",
					"price": "153.30",
					"price_open": "156.41",
					"day_high": "156.73",
					"day_low": "152.62",
					"52_week_high": "233.47",
					"52_week_low": "142.00",
					"day_change": "-3.52",
					"change_pct": "-2.24",
					"close_yesterday": "156.82",
					"market_cap": "725078814334",
					"volume": "30393970",
					"volume_avg": "45424373",
					"shares": "4729803000",
					"stock_exchange_long": "NASDAQ Stock Exchange",
					"stock_exchange_short": "NASDAQ",
					"timezone": "EST",
					"timezone_name": "America/New_York",
					"gmt_offset": "-18000",
					"last_trade_time": "2019-01-22 16:00:01"
				}]
	}`))
	}))
	defer positiveserv.Close()

	errorserv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
	}))
	defer errorserv.Close()

	type fields struct {
		client *httptest.Server
		host   string
	}
	type args struct {
		symbols   string
		exchanges string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string][]models.StockDetails
		wantErr bool
	}{
		{
			name:    "successfully retrieved stock details",
			fields:  fields{client: positiveserv, host: positiveserv.URL},
			args:    args{symbols: "AAPL,MSFT", exchanges: "NASDAQ"},
			want:    successresp,
			wantErr: false,
		},
		{
			name:    "failed to retrieve stock details",
			fields:  fields{client: errorserv, host: errorserv.URL},
			args:    args{symbols: "AAPL,MSFT", exchanges: ""},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stocks{
				client: tt.fields.client.Client(),
			}
			host = tt.fields.host
			got, err := s.GetStocks(tt.args.symbols, tt.args.exchanges)
			if (err != nil) != tt.wantErr {
				t.Errorf("stocks.Stocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stocks.Stocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStock(t *testing.T) {

	tests := []struct {
		name    string
		want    Stocks
		wantErr bool
	}{
		{
			name:    "failed to retrieved exact stock service object",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStock()
			if err != nil {
				t.Errorf("NewStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("NewStock() = %v, want %v", got, tt.want)
			}
		})
	}
}
