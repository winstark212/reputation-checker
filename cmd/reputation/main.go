package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
	feeds "github.com/winstark212/reputation-checker/pkg/feeds"
)

type ScanResult struct {
	CommonFeed map[string]string
}

var (
	feed = map[string]map[string]string{}
	scanResult ScanResult
)


func init() {
	ipFeedFile, err := ioutil.ReadFile("configs\\feeds.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ipFeedFile, &feed)
	if err != nil {
		log.Fatal(err)
	}
}

func updateFeed(feedType  string) {
	finflag := make(chan string)
	defer close(finflag)
	for k, v := range feed[feedType] {
		go feeds.UpdateFeedDB(v, k, finflag)
		<- finflag
	}
}

func searchIOC(keyword, feedType string) {
	for k, v := range feed[feedType] {
		
	}
}

func main() {

	feeds.GetAnalysisFromFile("", )
}