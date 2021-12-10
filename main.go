package main

import (
	"log"
	"strings"
	"time"

	"github.com/tomassirio/bitcoinTelegram/config"
	"github.com/tomassirio/bitcoinTelegram/handler"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  config.LoadConfig().Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for k, v := range handler.LoadHandler(b) {
		b.Handle(k, v)
		log.Println(k + "✅ Loaded!")
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Println(m.Text)
		if strings.Contains(strings.ToLower(m.Text), "crypto") {
			g := &tb.Animation{File: tb.FromURL("https://c.tenor.com/ndyV5-3mkisAAAAS/kissing-kiss.gif")}
			b.Send(m.Chat, "Someone said crypto???")
			b.Send(m.Chat, g)
		}
		// b.Send(m.Sender, "hello world")
	})
	log.Println("OnText ✅ Loaded!")

	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		log.Println("OnChannelPost: " + m.Text)
		if strings.Contains(strings.ToLower(m.Text), "crypto") {
			g := &tb.Animation{File: tb.FromURL("https://c.tenor.com/ndyV5-3mkisAAAAS/kissing-kiss.gif")}
			b.Send(m.Chat, "Someone said crypto???")
			b.Send(m.Chat, g)
		}
	})
	log.Println("OnChannelPost ✅ Loaded!")

	// blocks until shutdown
	b.Start()

}
