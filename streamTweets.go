package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
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
func isFromOwnBot(flock Botnet, tweet *twitter.Tweet) bool {
	for i := 0; i < len(flock.ScreenNames); i++ {
		if flock.ScreenNames[i] == tweet.User.ScreenName {
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
		fmt.Println(err)
	}
	if httpResp.Status != "200 OK" {
		c.Red("error: " + httpResp.Status)
		c.Purple("maybe twitter has blocked the account, CTRL+C, wait 15 minutes and try again")
	}
	fmt.Print("tweet posted: ")
	c.Green(tweet.Text)
}

func startStreaming(botnet Botnet, bot *twitter.Client, botScreenName string, keywords []string, replies []string) {
	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if isRT(tweet) == false && isFromOwnBot(botnet, tweet) == false {
			//processTweet(botnetUser, botScreenName, keywords, tweet)
			fmt.Println("[bot @" + botScreenName + "] - New tweet detected:")
			c.Yellow(tweet.Text)
			reply := getRandomReplyFromReplies(replies)
			fmt.Print("reply: ")
			c.Green(reply)
			fmt.Println(tweet.User.ScreenName)
			fmt.Println(tweet.ID)
			replyTweet(bot, "@"+tweet.User.ScreenName+" "+reply, tweet.ID)
			waitMinutes(1)
		}
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	fmt.Println("Starting Stream...")
	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         keywords,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := bot.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}
	// Receive messages until stopped or stream quits
	demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	/*ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()*/
}
func streamTweets(botnet Botnet, keywords []string, replies []string) {
	fmt.Println("total keywords: " + strconv.Itoa(len(keywords)))
	c.Purple("keywords to follow: ")
	fmt.Println(keywords)
	c.Green("Starting to stream tweets")
	for i := 0; i < len(botnet.Clients); i++ {
		go startStreaming(botnet, botnet.Clients[i], botnet.ScreenNames[i], keywords, replies)
		//wait 35 seconds to start the next bot
		waitSeconds(35)
	}
}
