package commands

import (
	"fmt"
	"math"
	"strconv"

	"cryptoTelegramBot/utils"
)

func GetSummary(sym string) (string, string, string, error) {
	p, err := utils.GetApiCall(sym)
	if p.Symbol == "" {
		return "", "", "", err
	}
	l, err := strconv.ParseFloat(p.Last, 32)
	o, err := strconv.ParseFloat(p.Open, 32)
	his := ((l - o) / o) * 100
	if !math.Signbit(float64(his)) {
		return p.Symbol, p.Last, "%" + fmt.Sprintf("%.2f", ((l-o)/o)*100), err
	} else {
		return p.Symbol, p.Last, "-%" + fmt.Sprintf("%.2f", -1*((l-o)/o)*100), err
	}
}
