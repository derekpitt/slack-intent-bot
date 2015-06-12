package slackintent

import "github.com/mholt/binding"

type SlackCommandData struct {
	Token       string
	ChannelID   string
	ChannelName string
	UserID      string
	UserName    string
	Text        string
}

func (s *SlackCommandData) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.Token:       "token",
		&s.ChannelID:   "channel_id",
		&s.ChannelName: "channel_name",
		&s.Text:        "text",
		&s.UserID:      "user_id",
		&s.UserName:    "user_name",
	}
}

func (d *SlackCommandData) GetUserToken() string {
	return "<@" + d.UserID + "|" + d.UserName + ">"
}

type incomingData struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}
