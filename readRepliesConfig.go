package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func readRepliesConfig() []string {
	var replies []string
	file, e := ioutil.ReadFile("repliesConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &replies)

	return replies
}
