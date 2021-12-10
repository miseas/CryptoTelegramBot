package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"cryptoTelegramBot/model"
)

func GetApiCall(sym string) (*model.Price, error) {
	// resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/ticker")
	var symbol string
	if sym == "" || len(sym) > 8 {
		symbol = "BTCEUR"
	} else {
		symbol = strings.ToUpper(sym)
	}

	resp, err := http.Get("https://api.binance.com/api/v3/ticker/24hr?symbol=" + symbol)
	p := &model.Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}
