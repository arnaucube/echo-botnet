package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func readKeywordsConfig() []string {
	var keywords []string
	file, e := ioutil.ReadFile("keywordsConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &keywords)

	return keywords
}
