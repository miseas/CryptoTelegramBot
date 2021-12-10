package handler

import (
	"log"

	"github.com/tomassirio/bitcoinTelegram/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		log.Println("PAYLOAD: " + m.Payload)
		symbol, price, _ := commands.GetPrice(m.Payload)
		if price == "" {
			b.Send(m.Chat, "Sorry I don't know "+m.Payload)
		} else {
			b.Send(m.Chat, symbol+"'s Current price is: "+price)
		}
	}

	commandMap["/historic"] = func(m *tb.Message) {
		log.Println("PAYLOAD: " + m.Payload)
		symbol, res, g, _ := commands.GetHistoric(m.Payload)
		if symbol == "" {
			b.Send(m.Chat, "Sorry I don't know "+m.Payload)
		} else {
			b.Send(m.Chat, symbol+"'s Price compared to yesterday is: "+res)
			b.Send(m.Chat, g)
		}
	}

	commandMap["/summary"] = func(m *tb.Message) {
		log.Println("PAYLOAD: " + m.Payload)
		symbol, p, h, _ := commands.GetSummary(m.Payload)
		if symbol == "" {
			b.Send(m.Chat, "Sorry I don't know "+m.Payload)
		} else {
			b.Send(m.Chat, symbol+"'s Current price is: "+p+"\n"+symbol+"'s Price compared to yesterday is: "+h)
		}
	}

	// commandMap["crypto"] = func(m *tb.Message) {
	// 	log.Println("PAYLOAD: " + m.Payload)
	// 	g := &tb.Animation{File: tb.FromURL("https://c.tenor.com/ndyV5-3mkisAAAAS/kissing-kiss.gif")}
	// 	b.Send(m.Chat, "Someone said crypto???")
	// 	b.Send(m.Chat, g)
	// }

	return commandMap
}
