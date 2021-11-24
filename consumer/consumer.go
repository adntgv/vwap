package consumer

import (
	"log"

	"github.com/adntgv/vwap/types"
)

type Consumer interface {
	Consume(chan types.VWAP)
}

type consumer struct{}

func NewConsumer() Consumer {
	return &consumer{}
}

func (c *consumer) Consume(vwaps chan types.VWAP) {
	for vwap := range vwaps {
		log.Println(vwap)
	}
}
