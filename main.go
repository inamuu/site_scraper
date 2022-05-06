package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
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

func main() {
	webhookUrl := os.Getenv("SLACKURL")
	err := slackNotify(webhookUrl)
	if err != nil {
		return
	}
	fmt.Println("ok")
}
