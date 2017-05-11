package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Config stores the data from json botsConfig.json file
type Config struct {
	Title             string `json:"title"`
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessTokenKey    string `json:"access_token_key"`
	AccessTokenSecret string `json:"access_token_secret"`
}

//Botnet stores each bot configured
type Botnet struct {
	ScreenNames []string
	Clients     []*twitter.Client
}

func readConfigTokensAndConnect() (botnet Botnet) {
	var config []Config
	var clients []*twitter.Client

	file, e := ioutil.ReadFile("botsConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
	fmt.Println("botnetConfig.json read comlete")

	fmt.Print("connecting to twitter api --> ")
	for i := 0; i < len(config); i++ {
		configu := oauth1.NewConfig(config[i].ConsumerKey, config[i].ConsumerSecret)
		token := oauth1.NewToken(config[i].AccessTokenKey, config[i].AccessTokenSecret)
		httpClient := configu.Client(oauth1.NoContext, token)
		// twitter client
		client := twitter.NewClient(httpClient)
		clients = append(clients, client)
		botnet.ScreenNames = append(botnet.ScreenNames, config[i].Title)
	}
	botnet.Clients = clients

	fmt.Println("connection successful")

	return botnet
}
