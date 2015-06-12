package slackintent

import (
	"net/http"

	"github.com/mholt/binding"
	"github.com/unrolled/render"
)

type Bot struct {
	OutgoingToken string
	IncomingURL   string

	handlers  map[string]IntentHandler
	witClient *witClient
	r         *render.Render
}

func NewBot(outGoingToken, incomingURL, witAPIKey string) *Bot {
	return &Bot{
		OutgoingToken: outGoingToken,
		IncomingURL:   incomingURL,

		handlers:  make(map[string]IntentHandler),
		witClient: newWitClient(witAPIKey),
		r:         render.New(),
	}
}

func (b *Bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := new(SlackCommandData)

	if err := binding.Bind(r, data); err != nil {
		b.r.Data(w, http.StatusBadRequest, []byte("bad request"))
		return
	}

	if data.Token != b.OutgoingToken {
		b.r.Data(w, http.StatusBadRequest, []byte("bad request"))
		return
	}

	// go off to witai
	outcomes, err := b.witClient.getOutcomes(data.Text)

	if err != nil {
		return
	}

	for _, outcome := range outcomes {
		if h, ok := b.handlers[outcome.Intent]; ok {
			c := &intentContext{
				botInstance: b,
				data:        data,
				outcome:     outcome,
			}

			h(c)
		}
	}
}

func (b *Bot) HandleIntent(intent string, handler IntentHandler) {
	b.handlers[intent] = handler
}
