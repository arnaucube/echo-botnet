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
	Title               string `json:"title"`
	Consumer_key        string `json:"consumer_key"`
	Consumer_secret     string `json:"consumer_secret"`
	Access_token_key    string `json:"access_token_key"`
	Access_token_secret string `json:"access_token_secret"`
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
		configu := oauth1.NewConfig(config[i].Consumer_key, config[i].Consumer_secret)
		token := oauth1.NewToken(config[i].Access_token_key, config[i].Access_token_secret)
		httpClient := configu.Client(oauth1.NoContext, token)
		// twitter client
		client := twitter.NewClient(httpClient)
		clients = append(clients, client)
		botnet.ScreenNames = append(botnet.ScreenNames, config[i].Title)
	}
	botnet.Clients = clients

	fmt.Println("connection successfull")

	return botnet
}
