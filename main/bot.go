package main

import (
	"os"
	"regexp"
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

func handleRequest(rtm *slack.RTM, ev *slack.MessageEvent) {
	info := rtm.GetInfo()
	text := ev.Text
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	match, _ := regexp.MatchString("!excuse", text)

	if ev.User != info.User.ID && match {
		excuse := GetExcuse()
		rtm.SendMessage(rtm.NewOutgoingMessage(excuse, ev.Channel, slack.RTMsgOptionTS(ev.ThreadTimestamp)))
		println("Sent: '" + excuse + "' to " + ev.Channel)
	}
}

func beBot(rtm *slack.RTM) {
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				handleRequest(rtm, ev)
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
