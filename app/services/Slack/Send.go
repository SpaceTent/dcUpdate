package Slack

import (
	"bytes"
	"encoding/json"
	l "log/slog"
	"net/http"

	"dcupdate/app/environment"
)

type SlackMessage struct {
	Channel  string  `json:"channel"`
	Username string  `json:"username"`
	Text     string  `json:"text"`
	Blocks   []Block `json:"blocks"`
}

type Block struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func QuickandDirtySend(channel string, message string) {
	slackURL := environment.GetEnvString("SLACK_URL", "") // Replace with your Slack URL
	if slackURL == "" {
		l.Error("No Slack URL")
		return
	}

	msg := SlackMessage{
		Channel:  "#" + channel,
		Username: "platform-notifications",
		Text:     "Deployment",
		Blocks: []Block{
			{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: message,
				},
			},
		},
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		// handle error
	}

	resp, err := http.Post(slackURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}
