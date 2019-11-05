package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
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
	log.Println("update ", feedType, " successfully!")
}

func searchIOC(feedType,keyword  string) {
	flag := make(chan string)
	for k, v := range feed[feedType] {
		go feeds.GetAnalysisFromFile(keyword, v, scanResult.CommonFeed, k, flag)
		<- flag
	}
	for k, v := range scanResult.CommonFeed{
		if v != "" {
			log.Println( "Found in ", k, " ", v)
		}
	}

}

func main() {
	var updateFeed = &cobra.Command{
		Use: "update [type of feed: domain, ip]",
		Short: "update commond feed, default now are update all",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				updateFeed(args[0])
			} else {
				updateFeed("domain")
				updateFeed("ip")
			}
		},
	}

	var cmdSearch = &cobra.Command{
		Use: "search [type of feed: domain, ip]",
		Short: "update commond feed, default now are update all",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				searchIOC(args[0], args[1])
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(updateFeed, cmdSearch)
	rootCmd.Execute()
}