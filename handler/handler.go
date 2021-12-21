package handler

import (
	"log"
	"strings"
	"time"

	"cryptoTelegramBot/commands"

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

	commandMap["/notification"] = func(m *tb.Message) {
		log.Println("PAYLOAD: " + m.Payload)
		b.Send(m.Chat, "Sorry this is WIP")
	}

	commandMap["/notification"] = func(m *tb.Message) {
		log.Println("PAYLOAD: " + m.Payload)
		slice := strings.Fields(m.Payload)
		switch slice[0] { // missing expression means "true"
		case "add":
			symbol, _ := commands.CreateNotification(m.Chat.Recipient(), slice[1], slice[2], slice[3])
			if symbol == "" {
				b.Send(m.Chat, "Sorry I don't know "+m.Payload)
			} else {
				b.Send(m.Chat, "Notification created for symbol: "+symbol)
			}
		case "remove":
			msg := commands.DeleteNotification(m.Chat.Recipient(), slice[1], slice[2], slice[3])
			b.Send(m.Chat, msg)
		case "list":
			msg := commands.GetNotificationsByUser(m.Chat.Recipient())
			b.Send(m.Chat, msg)
		default:
			b.Send(m.Chat, "I don't know the cmd: "+slice[0])
		}

	}

	return commandMap
}

func PullNotifications(channel chan string, b *tb.Bot) {
	log.Println("Start PullNotifications")
	for {
		select {
		// handle incoming updates
		default:
			// TODO: setup timer
			time.Sleep(60 * time.Second)
			notifications := commands.GetAllNotifications()
			for _, notif := range notifications {
				// log.Printf("\n----\nId: %d\nUserId: %s\nSymbol: %s\nCompare Value: %f\n----", notif.Id, notif.UserId, notif.Symbol, notif.CompareValue)
				alert_msg, _ := commands.CheckNotification(notif)
				if alert_msg != "" {
					client := Client{Client_id: notif.UserId}
					b.Send(client, alert_msg)
					if commands.UpdateNotificationCounter(notif) {
						b.Send(client, "Notification deleted after being notified 3 times")
					}
				}
			}
		}
	}
}

type Client struct {
	Client_id string
}

func (c Client) Recipient() string {
	return c.Client_id
}
