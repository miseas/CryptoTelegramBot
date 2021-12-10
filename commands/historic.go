package commands

import (
	"fmt"
	"math"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"

	"cryptoTelegramBot/utils"
)

func GetHistoric(sym string) (string, string, *tb.Animation, error) {
	p, err := utils.GetApiCall(sym)
	if p.Symbol == "" {
		return "", "", nil, err
	}
	l, err := strconv.ParseFloat(p.Last, 32)
	o, err := strconv.ParseFloat(p.Open, 32)
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		g := &tb.Animation{File: tb.FromURL("https://i.pinimg.com/originals/e4/38/99/e4389936b099672128c54d25c4560695.gif")}
		return p.Symbol, "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), g, err
	} else {
		g := &tb.Animation{File: tb.FromURL("http://www.brainlesstales.com/bitcoin-assets/images/fan-versions/2015-01-osEroUI.gif")}
		return p.Symbol, "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), g, err
	}
}
