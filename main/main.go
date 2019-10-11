package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

func main() {
	token := getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	go rtm.ManageConnection()
	go beBot(rtm)

	port := getenv("PORT")

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
