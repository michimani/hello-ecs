package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

const (
	ssmParamNameOfWebhook string = "/dev/slack-webhook"
	iconEmoji             string = ":mega:"
	userName              string = "Hello ECS Job"
)

type payload struct {
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	UserName  string `json:"username"`
}

func getWebhookURL() (string, error) {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := ssm.NewFromConfig(config)
	out, err := client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String(ssmParamNameOfWebhook),
	})

	if err != nil {
		return "", err
	}

	webhookURL := aws.ToString(out.Parameter.Value)
	return webhookURL, nil
}

func postToSlack(message, webhookURL string) error {
	p, err := json.Marshal(payload{
		Text:      message,
		IconEmoji: iconEmoji,
		UserName:  userName,
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
		fmt.Println("Failed to get Slack Webhook URL.")
		return
	}

	if err := postToSlack(*m, webhookUrl); err != nil {
		fmt.Println("Failed to post message to Slack.")
	}
}
