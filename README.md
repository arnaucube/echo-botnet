http://arnaucode.com/echo-botnet/

# echo-botnet [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/echo-botnet)](https://goreportcard.com/report/github.com/arnaucode/echo-botnet)
A twitter botnet with autonomous bots replying tweets with pre-configured replies


Echo (Ēkhō): https://en.wikipedia.org/wiki/Echo_(mythology)#


# How it works
- each bot is launched every 35 seconds, so the bot1 is launched at time 0, the bot2 is launched at time 35s, the bot3 is launched at time 1:10 minutes, etc
- the bots are running in parallel
- each bot streams the Twitter API getting the tweets containing one or more keywords configured in the file keywordsConfig.json
- then, the bot randomly gets a reply from the file repliesConfig.json
- after a bot replies a tweet, the bot 'sleeps' 1 minute, to avoid the Twitter API limitation

Needs the files:
```
botsConfig.json     --> contains the tokens of the bots
keywordsConfig.json --> contains the keywords to track from Twitter API
repliesConfig.json  --> contains the replies
```

Just need to configure (deleting the word 'DEMO' from the filenames):

![echo-botnet](https://raw.githubusercontent.com/arnaucode/echo-botnet/master/configurationScreenshot.png "configurationScreenshot")


To run, on the /build directory:
```
./echo-botnet
```
Also, can run from the sources, with the Go compiler installed. On the .go files directory:
```
go run *.go
```
