package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"github.com/joho/godotenv"
)

type payload struct {
	Text string `json:"text"`
}

func slackNotify(webhookUrl string) (err error) {
	jsonP, err := json.Marshal(payload{Text: "Hello, World!"})
	if err != nil {
		return err
	}

	resp, err := http.PostForm(
		webhookUrl,
		url.Values{
			"payload": {string(jsonP)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return
}

func requestUrl(url string) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	return
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	url := os.Getenv("REQUESTURL")
	err = requestUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	//webhookUrl := os.Getenv("SLACKURL")
	//err := slackNotify(webhookUrl)
	//if err != nil {
	//	fmt.Println("ng")
	//	return
	//}
	//fmt.Println("ok")
}
