package parser

import (
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	testCase := `{"type":"match","trade_id":241016739,"maker_order_id":"50712110-7fb2-4229-83cb-7fe3268d853b","taker_order_id":"d3571d7b-0df7-4ac3-91ff-5d9179dacfe1","side":"sell","size":"0.0008414","price":"56713.13","product_id":"BTC-USD","sequence":31412488430,"time":"2021-11-24T16:11:10.255620Z"}`

	p := NewParser()

	transaction, err := p.Parse(strings.NewReader(testCase))
	if err != nil {
		t.Error(err)
	}

	if transaction.Price != 56713.13 {
		t.Errorf("expected %v, got %v", transaction.Price, 56713.13)
	}

	if transaction.ProductID != "BTC-USD" {
		t.Errorf("expected %v, got %v", transaction.ProductID, "BTC-USD")
	}

	if transaction.Size != 0.0008414 {
		t.Errorf("expected %v, got %v", transaction.Size, 0.0008414)
	}
}
