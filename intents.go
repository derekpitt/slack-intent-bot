package slackintent

type IntentContext interface {
	SlackData() *SlackCommandData
	WitOutcome() outcome

	Reply(text string) error
}

type IntentHandler func(ctx IntentContext)
