package lib

import (
	"os"

	"github.com/slack-go/slack"
)

// SlackMessage is a struct for sending a message to Slack
type SlackMessage struct {
	Channel   string
	Username  string
	Text      string
	IconEmoji string
}

var WEBHOOK_URL = os.Getenv("WEBHOOK_URL")

// SendSlackMessage sends a message to Slack using the Webhook URL
func SendSlackMessage(sm SlackMessage) error {
	// attachment := slack.Attachment{
	// 	AuthorName:    "notifyCreditHistory",
	// 	// AuthorIcon:    "https://avatars2.githubusercontent.com/u/652790",
	// 	Text:          sm.Text,
	// }
	msg := slack.WebhookMessage{
		// Attachments: []slack.Attachment{attachment},
		Username: sm.Username,
		IconEmoji: sm.IconEmoji,
		Text: sm.Text,
	}

	if err := slack.PostWebhook(WEBHOOK_URL, &msg); err != nil {
		return err
	}
	return nil
}