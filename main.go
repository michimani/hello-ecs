package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	iconEmoji string = ":mega:"
	userName  string = "Hello ECS Job"
	channel   string = "dev"
)

type payload struct {
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	UserName  string `json:"username"`
	Channel   string `json:"channel"`
}

func getWebhookURL() (string, error) {
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		return "", errors.New("The environment value SLACK_WEBHOOK_URL is required.")
	}
	return webhookURL, nil
}

func postToSlack(message, webhookURL string) error {
	p, err := json.Marshal(payload{
		Text:      message,
		IconEmoji: iconEmoji,
		UserName:  userName,
		Channel:   channel,
	})
	if err != nil {
		return err
	}
	resp, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	m := flag.String("m", "Hello ECS!!!", "Message posted to Slack")
	flag.Parse()

	webhookUrl, err := getWebhookURL()
	if err != nil {
		fmt.Println("Failed to get Slack Webhook URL.", err.Error())
		return
	}

	if err := postToSlack(*m, webhookUrl); err != nil {
		fmt.Println("Failed to post message to Slack.", err.Error())
	}
}
