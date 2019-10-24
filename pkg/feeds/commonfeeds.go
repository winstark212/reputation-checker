package feeds

import (
	"time"
	"log"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
)

// GetAnalysisFromFile search ioc against source files
func GetAnalysisFromFile(search, url string, result map[string]string, key string, flag chan string) {
	filePath := "pkg/feeds/db/" + key
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	res := string(fileBytes)
	
	if strings.Contains(res, search) {
		result[key] = "Malicious"
	} else {
		result[key] = ""
	}
	flag <- "done"
}

// UpdateFeedDB update feed database
func UpdateFeedDB(url string, key string, finflag chan string) {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	dataBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	filePath := "pkg/feeds/db/" + key
	op, err := os.Create(filePath)
	defer op.Close()
	if err != nil {
		log.Println("Error while updating feed", err)
	}
	op.Write(dataBody)
	finflag <- "done"
}