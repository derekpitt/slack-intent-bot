package slackintent

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/wealth-ai/go-wit"
)

type intentContext struct {
	data        *SlackCommandData
	outcome     wit.Outcome
	botInstance *Bot
}

func (c *intentContext) SlackData() *SlackCommandData {
	return c.data
}

func (c *intentContext) WitOutcome() wit.Outcome {
	return c.outcome
}

func (c *intentContext) Reply(text string) error {
	d, _ := json.Marshal(incomingData{
		Text:    text,
		Channel: "#" + c.data.ChannelName,
	})

	resp, err := http.Post(c.botInstance.IncomingURL, "application/javascript", bytes.NewReader(d))

	if err == nil {
		defer resp.Body.Close()
	}

	return err
}
