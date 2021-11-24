package calculator

import (
	"github.com/adntgv/vwap/buffer"
	"github.com/adntgv/vwap/types"
)

type Calculator interface {
	Process(in chan *types.Transaction)
}

type distributedCalculator struct {
	calculators         []*calculator
	transactionChannels map[string]chan *types.Transaction
}

func NewDisctributedCalculator(size int, productIDs []string, out chan types.VWAP) Calculator {
	calculators := make([]*calculator, len(productIDs))
	transactionChannels := make(map[string]chan *types.Transaction)

	for i, productID := range productIDs {
		in := make(chan *types.Transaction)
		calculators[i] = newCalculator(size, in, out)
		transactionChannels[productID] = in
	}

	return &distributedCalculator{
		calculators:         calculators,
		transactionChannels: transactionChannels,
	}
}

func (d *distributedCalculator) Process(in chan *types.Transaction) {
	for _, calculator := range d.calculators {
		go calculator.calculate()
	}

	for t := range in {
		transactionChannel, ok := d.transactionChannels[t.ProductID]
		if ok {
			transactionChannel <- t
		}
	}
}

type calculator struct {
	buf            *buffer.RingBuffer
	totalSizePrice float64
	totalSize      float64
	in             chan *types.Transaction
	out            chan types.VWAP
}

func newCalculator(size int, in chan *types.Transaction, out chan types.VWAP) *calculator {
	return &calculator{
		buf: buffer.NewRingBuffer(size),
		in:  in,
		out: out,
	}
}

func (c *calculator) calculate() {
	for t := range c.in {
		c.out <- c.process(t)
	}
}

func (c *calculator) process(t *types.Transaction) types.VWAP {
	sizePrice := t.Size * t.Price
	fresh := types.BufferedType{
		SizePrice: sizePrice,
		Size:      t.Size,
	}

	c.add(fresh)
	old := c.buf.Add(fresh)
	c.subtract(old)

	return c.vwap(t.ProductID)
}

func (c *calculator) add(t types.BufferedType) {
	c.totalSizePrice += t.SizePrice
	c.totalSize += t.Size
}

func (c *calculator) subtract(t types.BufferedType) {
	c.totalSizePrice -= t.SizePrice
	c.totalSize -= t.Size
}

func (c *calculator) vwap(productID string) types.VWAP {
	return types.VWAP{
		Value:     c.totalSizePrice / c.totalSize,
		ProductID: productID,
	}
}
