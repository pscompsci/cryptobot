package main

import (
	"fmt"

	ws "github.com/aopoltorzhicky/go_kraken/websocket"
	"github.com/pscompsci/cryptobot/internal/exchanges"
	"github.com/pscompsci/cryptobot/internal/server"
	"github.com/pscompsci/cryptobot/pkg/events"
)

func main() {
	eb := events.NewEventBus()
	krakenChannel := make(chan events.DataEvent)

	eb.Subscribe("kraken", krakenChannel)

	kraken := exchanges.NewKraken("", "", eb)
	server := server.New(eb)

	go server.Serve()
	go kraken.Listen([]string{ws.BTCUSD, ws.ETHUSD, ws.ADAUSD, ws.DASHUSD, ws.LTCUSD, ws.EOSUSD})

	for {
		select {
		case update := <-krakenChannel:
			packet := update.Data.(ws.Update)
			switch data := packet.Data.(type) {
			case ws.TickerUpdate:
				fmt.Printf("%s - Ask: %s, Volume: %s\n", packet.Pair, data.Ask.Price.String(), data.Ask.Volume.String())
			}
		}
	}
}
