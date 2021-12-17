package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"cryptoTelegramBot/config"
	"cryptoTelegramBot/handler"
	"cryptoTelegramBot/repo"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	repo.Start_db()
	repo.Create_tables()

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
			log.Printf("struct1: %v\n", m.Chat)
			b.Send(m.Chat, "Someone said crypto???")
			b.Send(m.Chat, g)
		}
		// b.Send(m.Sender, "hello world")
	})
	log.Println("OnText ✅ Loaded!")

	// channel := make(chan string)
	// go test_async(channel)

	// for {
	// 	select {
	// 	// handle incoming updates
	// 	case upd := <-channel:
	// 		log.Println("salimos con data" + upd)
	// 		// call to stop polling
	// 		// case <-b.stop:
	// 		// 	close(stop)
	// 		// 	return
	// 	}
	// }

	// client := Client{client_id: "369774783"}
	// log.Printf("struct1: %v\n", client)
	// b.Send(client, "hola capo!")

	func1 := func() {
		// blocks until shutdown
		b.Start()
	}

	channel := make(chan string)
	func2 := func() {
		handler.PullNotifications(channel, b)
	}

	Parallelize(func1, func2)

}

func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}
