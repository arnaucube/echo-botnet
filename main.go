package main

import "fmt"

const version = "0.1-dev"

func main() {
	c.Yellow("echo-botnet")
	fmt.Println("---------------")
	c.Cyan("echo-botnet initialized")
	c.Purple("https://github.com/arnaucode/echo-botnet")
	fmt.Println("version " + version)

	keywords := readKeywordsConfig()
	c.Cyan("keywords configured: ")
	fmt.Println(keywords)
	replies := readRepliesConfig()
	c.Cyan("replies configured: ")
	fmt.Println(replies)

	fmt.Println("Reading botsConfig.json file")
	botnet := readConfigTokensAndConnect()
	c.Cyan("[list of configured bots]:")
	for _, v := range botnet.ScreenNames {
		c.Cyan(v)
	}
	streamTweets(botnet, keywords, replies)
}
