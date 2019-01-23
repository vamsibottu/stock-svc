package models

// StockDetails is the structure of the response data of a specific symbol
type StockDetails struct {
	Symbol         string `json:"symbol,omitempty"`
	Name           string `json:"name,omitempty"`
	Price          string `json:"price,omitempty"`
	CloseYesterday string `json:"close_yesterday,omitempty"`
	MarketCap      string `json:"market_cap,omitempty"`
	Volume         string `json:"volume,omitempty"`
	TimeZone       string `json:"timezone,omitempty"`
	TimeZoneName   string `json:"timezone_name,omitempty"`
	GmtOffset      string `json:"gmt_Offset,omitempty"`
	LastTradeTime  string `json:"last_trade_time,omitempty"`
}
