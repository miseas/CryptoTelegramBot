package main

import (
	"log"
	"strings"
	"time"

	"cryptoTelegramBot/config"
	"cryptoTelegramBot/handler"
	"cryptoTelegramBot/repo"
	"cryptoTelegramBot/utils"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	repo.Start_db()
	repo.Create_tables()

	b, err := tb.NewBot(tb.Settings{
		// You can set custom API URL. If empty it equals to "https://api.telegram.org".
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
			g := &tb.Animation{File: tb.FromURL("https://c.tenor.com/hJ834RcJZbYAAAAC/elon-musk-dance.gif")}
			log.Printf("struct1: %v\n", m.Chat)
			b.Send(m.Chat, "Someone said crypto???")
			b.Send(m.Chat, g)
		}
	})
	log.Println("OnText ✅ Loaded!")

	func1 := func() {
		// blocks until shutdown
		b.Start()
	}

	channel := make(chan string)
	func2 := func() {
		handler.PullNotifications(channel, b)
	}

	utils.Parallelize(func1, func2)

}
