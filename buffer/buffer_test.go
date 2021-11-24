package buffer_test

import (
	"testing"

	"github.com/adntgv/vwap/buffer"
	"github.com/adntgv/vwap/types"
)

func TestRingBuffer_Add(t *testing.T) {
	buf := buffer.NewRingBuffer(1)

	old := buf.Add(types.BufferedType{
		SizePrice: 1,
		Size:      2,
	})
	if old.Size != 0 {
		t.Error("size is not zero")
	}
	if old.SizePrice != 0 {
		t.Error("size price is not zero")
	}

	old = buf.Add(types.BufferedType{
		SizePrice: 3,
		Size:      4,
	})
	if old.Size != 2 {
		t.Error("size is not 2")
	}
	if old.SizePrice != 1 {
		t.Error("size price is not 1")
	}
}
