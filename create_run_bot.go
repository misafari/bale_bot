package bale_bot

import "net/http"

type Bot struct {
	Token  string `json:"token"`
	Debug  bool   `json:"debug"`
	Buffer int    `json:"buffer"`

	//Self            User         `json:"-"`
	Client          *http.Client `json:"-"`
	shutdownChannel chan interface{}

	apiEndpoint string
}

func NewBot(token string) (*Bot, error) {
	panic("not implement")
}

func NewBotAPIWithClient(token, apiEndpoint string, client *http.Client) (*Bot, error) {
	bot := &Bot{
		Token:           token,
		Client:          client,
		Buffer:          100,
		shutdownChannel: make(chan interface{}),

		apiEndpoint: apiEndpoint,
	}

	//self, err := bot.GetMe()
	//if err != nil {
	//	return nil, err
	//}
	//
	//bot.Self = self

	return bot, nil
}
