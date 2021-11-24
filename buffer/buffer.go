package buffer

import "github.com/adntgv/vwap/types"

type RingBuffer struct {
	buf []types.BufferedType
	p   int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buf: make([]types.BufferedType, size),
	}
}

func (b *RingBuffer) Add(v types.BufferedType) types.BufferedType {
	if b.p == len(b.buf) {
		b.p = 0
	}

	tmp := b.buf[b.p]
	b.buf[b.p] = v
	b.p++

	return tmp
}
