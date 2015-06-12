package slackintent

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type witClient struct {
	apiKey string
}

func newWitClient(apiKey string) *witClient {
	return &witClient{
		apiKey: apiKey,
	}
}

type messageResponse struct {
	Text     string    `json:"_text"`
	Outcomes []outcome `json:"outcomes"`
}

type outcome struct {
	Text       string                     `json:"_text"`
	Intent     string                     `json:"intent"`
	IntentId   string                     `json:"intent_id"`
	Entities   map[string][]messageEntity `json:"entities"`
	Confidence float32                    `json:"confidence"`
}

type messageEntity struct {
	Metadata *string        `json:"metadata,omitempty"`
	Value    *interface{}   `json:"value,omitempty"`
	Grain    *string        `json:"grain,omitempty"`
	Type     *string        `json:"type,omitempty"`
	Unit     *string        `json:"unit,omitempty"`
	Body     *string        `json:"body,omitempty"`
	Entity   *string        `json:"entity,omitempty"`
	Start    *int64         `json:"start,omitempty"`
	End      *int64         `json:"end,omitempty"`
	Values   *[]interface{} `json:"values,omitempty"`
}

func (w *witClient) getOutcomes(text string) ([]outcome, error) {
	u, err := url.Parse("https://api.wit.ai/message?v=20150612")

	if err != nil {
		return []outcome{}, err
	}

	qs := u.Query()
	qs.Add("q", text)

	u.RawQuery = qs.Encode()

	client := &http.Client{}

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return []outcome{}, err
	}

	req.Header.Add("Authorization", "Bearer "+w.apiKey)

	resp, err := client.Do(req)

	if err != nil {
		return []outcome{}, err
	}

	defer resp.Body.Close()

	var message messageResponse

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&message)

	if err != nil {
		return []outcome{}, err
	}

	return message.Outcomes, nil
}
