package model

type Price struct {
	Last            string `json:"lastPrice"`
	PriceBeforeLast string `json:"prevClosePrice"`
	Open            string `json:"openPrice"`
	High            string `json:"highPrice"`
	Low             string `json:"lowPrice"`
	Vwap            string `json:"weightedAvgPrice"`
	Volume          string `json:"volume"`
	Bid             string `json:"bidPrice"`
	Ask             string `json:"askPrice"`
	Symbol          string `json:"symbol"`
}

// https://bitex.la/developers#api_ticker
/*
{
  "last":               1230.0,  // Last transaction price
  "price_before_last":  1220.0,  // Helps you tell if price is going up or down.
  "open":        1198.45875559,  // What the price was 24 hours ago.
  "high":               1230.0,  // Highest price for the past 24 hours.
  "low":          1193.2507548,  // Lowest price for the past 24 hours.
  "vwap":        1208.57944642,  // Volume-Weighted Average Price for the past 24 hours.
  "volume":        16.45315074,  // Transacted volume for the last 24 hours.
  "bid":          1226.5583985,  // Highest current buy order.
  "ask":         1235.71481927   // Lowest current ask order.
}
*/
