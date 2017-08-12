package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Bot stores the data from json botsConfig.json file
type Bot struct {
	Title             string `json:"title"`
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessTokenKey    string `json:"access_token_key"`
	AccessTokenSecret string `json:"access_token_secret"`
	Client            *twitter.Client
	SinceTweeted      int64 `json:"sincetweeted"`
	Blocked           bool  `json:"blocked"`
	SinceBlocked      int64 `json:"sinceblocked"`
}

var botnet []Bot
var keywords []string
var replies []string

func readConfigTokensAndConnect() {
	file, err := ioutil.ReadFile("botsConfig.json")
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &botnet)
	log.Println("botnetConfig.json read comlete")
	log.Print("connecting to twitter api")
	for k, _ := range botnet {
		configu := oauth1.NewConfig(botnet[k].ConsumerKey, botnet[k].ConsumerSecret)
		token := oauth1.NewToken(botnet[k].AccessTokenKey, botnet[k].AccessTokenSecret)
		httpClient := configu.Client(oauth1.NoContext, token)
		// twitter client
		client := twitter.NewClient(httpClient)
		botnet[k].Client = client
		botnet[k].Blocked = false
	}

	log.Println("connection successful")
}

func readKeywordsConfig() {
	file, err := ioutil.ReadFile("keywordsConfig.json")
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &keywords)
}

func readRepliesConfig() {
	file, err := ioutil.ReadFile("repliesConfig.json")
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &replies)
}
