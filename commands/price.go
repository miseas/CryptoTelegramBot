package commands

import (
	"log"

	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetPrice(sym string) (string, string, error) {
	log.Println("Price!")
	p, err := utils.GetApiCall(sym)
	log.Println(p.Last)
	return p.Symbol, p.Last, err
}
