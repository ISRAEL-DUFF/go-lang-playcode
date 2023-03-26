package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("USER_OAUTH_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Pretext: "Super Bot Message from Duff",
		Text:    "This is a message from Israel Duff.",
		Color:   "4af030",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, timestamp, err := client.PostMessage(channelId, slack.MsgOptionAttachments(attachment))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Message sent at %s", timestamp)
}
