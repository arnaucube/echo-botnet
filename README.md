# echo-botnet [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/echo-botnet)](https://goreportcard.com/report/github.com/arnaucode/echo-botnet)
A twitter botnet with autonomous bots replying tweets with pre-configured replies


Echo (Ēkhō): https://en.wikipedia.org/wiki/Echo_(mythology)#


Needs the files:
```
botsConfig.json     --> contains the tokens of the bots
keywordsConfig.json --> contains the keywords to track from Twitter API
repliesConfig.json  --> contains the replies
```

configuration file example (botsConfig.json):
```
[{
        "title": "account1",
        "consumer_key": "xxxxxxxxxxxxx",
        "consumer_secret": "xxxxxxxxxxxxx",
        "access_token_key": "xxxxxxxxxxxxx",
        "access_token_secret": "xxxxxxxxxxxxx"
    },
    {
        "title": "account2",
        "consumer_key": "xxxxxxxxxxxxx",
        "consumer_secret": "xxxxxxxxxxxxx",
        "access_token_key": "xxxxxxxxxxxxx",
        "access_token_secret": "xxxxxxxxxxxxx"
    },
    {
        "title": "account3",
        "consumer_key": "xxxxxxxxxxxxx",
        "consumer_secret": "xxxxxxxxxxxxx",
        "access_token_key": "xxxxxxxxxxxxx",
        "access_token_secret": "xxxxxxxxxxxxx"
    }
]

```
