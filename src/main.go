package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

type feedURL struct {
	Feed map[string]string
}

var (
	feed feedURL
)

func init() {
	ipFeedFile, err := ioutil.ReadFile("ipfeeds.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(ipFeedFile, &feed)
}

func main() {
	log.Println("halo")
}