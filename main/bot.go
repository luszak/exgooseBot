package main

import (
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("Missing envvar: " + name)
	}
	return v
}

func beBot(rtm *slack.RTM) {
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			println("Event received")
			println(msg.Type)
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				info := rtm.GetInfo()
				text := ev.Text
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)

				if ev.User != info.User.ID && text == "!excuse" {
					excuse := GetExcuse()
					rtm.SendMessage(rtm.NewOutgoingMessage(excuse, ev.Channel))
				}
			case *slack.RTMError:
				print("Got slack error: " + ev.Error())
			case *slack.InvalidAuthEvent:
				print("Invalid credentials")
				break Loop
			default:
			}
		}
	}
}
