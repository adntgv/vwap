package main

import (
	"flag"
	"log"
	"strings"

	"github.com/adntgv/vwap/calculator"
	"github.com/adntgv/vwap/consumer"
	"github.com/adntgv/vwap/parser"
	"github.com/adntgv/vwap/streamer"
	"github.com/adntgv/vwap/types"
)

var (
	products   = flag.String("product_ids", "BTC-USD,ETH-USD,ETH-BTC", "coma separated list of product ids")
	windowSize = flag.Int("size", 200, "window size")
)

func main() {
	flag.Parse()

	vwaps := make(chan types.VWAP)
	productIDs := strings.Split(*products, ",")
	transactions := make(chan *types.Transaction)

	consumer := consumer.NewConsumer()
	calculators := calculator.NewDisctributedCalculator(*windowSize, productIDs, vwaps)

	streamer, err := streamer.NewStreamer(productIDs, parser.NewParser())
	if err != nil {
		log.Fatalln(err)
	}

	go streamer.Stream(transactions)
	go calculators.Process(transactions)

	consumer.Consume(vwaps)
}
