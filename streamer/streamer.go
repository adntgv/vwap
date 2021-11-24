package streamer

import (
	"log"

	"github.com/adntgv/vwap/parser"
	"github.com/adntgv/vwap/types"
	"github.com/gorilla/websocket"
)

const (
	coinbaseAddress = "wss://ws-feed.exchange.coinbase.com"
)

type request struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

type Streamer interface {
	Stream(out chan *types.Transaction)
}

type coinbaseStreamer struct {
	c      *websocket.Conn
	parser parser.Parser
}

func NewStreamer(productIDs []string, parser parser.Parser) (Streamer, error) {
	c, _, err := websocket.DefaultDialer.Dial(coinbaseAddress, nil)
	if err != nil {
		return nil, err
	}

	streamer := &coinbaseStreamer{
		c:      c,
		parser: parser,
	}

	err = streamer.init(productIDs)
	if err != nil {
		return nil, err
	}

	return streamer, nil
}

func (streamer *coinbaseStreamer) init(productIDs []string) error {
	return streamer.c.WriteJSON(request{
		Type:       "subscribe",
		ProductIDs: productIDs,
		Channels:   []string{"matches"},
	})
}

func (streamer *coinbaseStreamer) Stream(out chan *types.Transaction) {
	for {
		mt, r, err := streamer.c.NextReader()
		if err != nil {
			log.Println(err)
			continue
		}

		var t *types.Transaction
		switch mt {
		case websocket.TextMessage:
			if t, err = streamer.parser.Parse(r); err != nil {
				log.Println(err)
				continue
			}
		default:
			continue
		}

		out <- t
	}
}
