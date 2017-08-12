package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
)

func isRT(tweet *twitter.Tweet) bool {
	tweetWords := strings.Split(tweet.Text, " ")
	for i := 0; i < len(tweetWords); i++ {
		if tweetWords[i] == "RT" {
			return true
		}
	}
	return false
}
func isFromBotnet(tweet *twitter.Tweet) bool {
	for i := 0; i < len(botnet); i++ {
		if botnet[i].Title == tweet.User.ScreenName {
			return true
		}
	}
	return false
}

func getRandomReplyFromReplies(replies []string) string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(replies))
	return replies[random]
}

func replyTweet(client *twitter.Client, text string, inReplyToStatusID int64) {
	tweet, httpResp, err := client.Statuses.Update(text, &twitter.StatusUpdateParams{
		InReplyToStatusID: inReplyToStatusID,
	})
	if err != nil {
		log.Println(err)
	}
	if httpResp.Status != "200 OK" {
		color.Red("error: " + httpResp.Status)
		log.Println("error" + httpResp.Status)
		color.Cyan("maybe twitter has blocked the account, CTRL+C, wait 15 minutes and try again")
		log.Println("maybe twitter has blocked the account, CTRL+C, wait 15 minutes and try again")
	}
	log.Println("tweet posted: " + tweet.Text)
}

func startStreaming(bot Bot) {
	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if isRT(tweet) == false && isFromBotnet(tweet) == false {
			//processTweet(botnetUser, botScreenName, keywords, tweet)
			log.Println("[bot @" + bot.Title + "] - New tweet detected:")
			log.Println(tweet.Text)
			reply := getRandomReplyFromReplies(replies)
			log.Println("reply: " + reply + ", to: @" + tweet.User.ScreenName /* + ". tweet ID: " + tweet.ID*/)

			//replyTweet(bot, "@"+tweet.User.ScreenName+" "+reply, tweet.ID)
			color.Green("replying tweet!")
			bot.SinceTweeted = time.Now().Unix()
		}
	}

	fmt.Println("Starting Stream...")
	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         keywords,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := bot.Client.Streams.Filter(filterParams)
	check(err)
	// Receive messages until stopped or stream quits
	demux.HandleChan(stream.Messages)
	/*for message := range stream.Messages {
		demux.Handle(message)
		log.Println("stopping stream")
		stream.Stop()
	}*/
}
