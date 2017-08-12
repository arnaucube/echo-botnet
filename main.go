package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

const version = "0.1-dev"

func main() {
	savelog()
	color.Blue("echo-botnet")
	fmt.Println("---------------")
	color.Cyan("echo-botnet initialized")
	color.Green("https://github.com/arnaucode/echo-botnet")
	log.Println("version " + version)

	readKeywordsConfig()
	color.Cyan("keywords configured: ")
	fmt.Println(keywords)
	readRepliesConfig()
	color.Cyan("replies configured: ")
	fmt.Println(replies)

	fmt.Println("Reading botsConfig.json file")
	readConfigTokensAndConnect()
	color.Cyan("[list of configured bots]:")
	for _, bot := range botnet {
		color.Cyan(bot.Title)
	}
	var i int
	for {
		sinceTweeted := time.Unix(botnet[i].SinceTweeted, 0)
		elapsed := time.Since(sinceTweeted)
		if elapsed.Seconds() > 60 {
			color.Blue("starting to use bot: " + botnet[i].Title)
			startStreaming(&botnet[i])
		} else {
			//log.Println("bot: " + botnet[i].Title + " not used due bot is in waiting time")
		}
		i++
		if i > len(botnet)-1 {
			i = 0
		}
	}
}
