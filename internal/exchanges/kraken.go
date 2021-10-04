package exchanges

import (
	"github.com/pscompsci/cryptobot/pkg/events"

	ws "github.com/aopoltorzhicky/go_kraken/websocket"
)

type kraken struct {
	apiKey    string
	secretKey string
	eventBus  *events.EventBus
}

func NewKraken(api, secret string, eb *events.EventBus) *kraken {
	return &kraken{
		apiKey:    api,
		secretKey: secret,
		eventBus:  eb,
	}
}

func (k *kraken) Listen(tickers []string) error {
	kraken := ws.NewKraken(ws.ProdBaseURL)
	if err := kraken.Connect(); err != nil {
		return err
	}
	if err := kraken.SubscribeTicker(tickers); err != nil {
		return err
	}

	for {
		select {
		case update := <-kraken.Listen():
			k.eventBus.Publish("kraken", update)
		}
	}
}
