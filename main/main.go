package main

import "github.com/nlopes/slack"

func main() {
	token := getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	go rtm.ManageConnection()
	beBot(rtm)
}
