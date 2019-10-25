package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
	feeds "github.com/winstark212/reputation-checker/pkg/feeds"
)

// ScanResult scan result against feeds
type ScanResult struct {
	CommonFeed map[string]string
}

var (
	feed = map[string]map[string]string{}
	scanResult ScanResult
)


func init() {
	ipFeedFile, err := ioutil.ReadFile("configs\\feeds.json") // use configs/feeds.json for linux-based
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ipFeedFile, &feed)
	if err != nil {
		log.Fatal(err)
	}
	scanResult.CommonFeed = make(map[string]string)
}

func updateFeed(feedType  string) {
	flag := make(chan string)
	defer close(flag)
	for k, v := range feed[feedType] {
		go feeds.UpdateFeedDB(v, k, flag)
		<- flag
	}
}

func searchIOC(keyword, feedType string) {
	flag := make(chan string)
	for k, v := range feed[feedType] {
		go feeds.GetAnalysisFromFile(keyword, v, scanResult.CommonFeed, k, flag)
		<- flag
	}
}

func main() {
	// updateFeed("domain")
	searchIOC("avvmail.com", "domain")
	// searchIOC("23.94.213.222", "ip")
	log.Println(scanResult.CommonFeed)
	log.Println("halo")
}