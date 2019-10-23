package feeds

import (
	"time"
	"log"
	"os"
	"net/http"
	"io/ioutil"
)

// UpdateFeedDB update feed database
func UpdateFeedDB(url string, key string) {
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
	filePath := "feeds/db/" + key
	op, err := os.Create(filePath)
	defer op.Close()
	if err != nil {
		log.Println("Error while updating feed")
	}
	op.Write(dataBody)
}