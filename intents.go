package slackintent

import "github.com/wealth-ai/go-wit"

type IntentContext interface {
	SlackData() *SlackCommandData
	WitOutcome() wit.Outcome

	Reply(text string) error
}

type IntentHandler func(ctx IntentContext)
