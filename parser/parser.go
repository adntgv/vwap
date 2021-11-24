package parser

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"

	"github.com/adntgv/vwap/types"
)

type Parser interface {
	Parse(r io.Reader) (*types.Transaction, error)
}

type parser struct {
	s *scanner.Scanner
}

func NewParser() Parser {
	return &parser{
		s: &scanner.Scanner{},
	}
}

func (p *parser) Parse(r io.Reader) (out *types.Transaction, err error) {
	p.s.Init(r)

	return parse(p.s)
}

func parse(s *scanner.Scanner) (t *types.Transaction, err error) {
	t = new(types.Transaction)
	defined := map[string]bool{
		`"product_id"`: true,
		`"size"`:       true,
		`"price"`:      true,
	}

	var expect string

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		text := s.TokenText()
		if text == `:` || text == "," || text == "}" {
			continue
		} else if defined[text] {
			expect = text
			continue
		} else {
			if len(text) > 1 {
				text = text[1 : len(text)-1] // quick trim
			} else {
				continue
			}

			switch expect {
			case `"product_id"`:
				t.ProductID = text
				expect = ""
			case `"size"`:
				t.Size, err = strconv.ParseFloat(text, 64)
				if err != nil {
					return nil, err
				}
				expect = ""
			case `"price"`:
				t.Price, err = strconv.ParseFloat(text, 64)
				if err != nil {
					return nil, err
				}
				expect = ""
			default:
				continue
			}
		}
	}

	if t.ProductID == "" {
		return nil, fmt.Errorf("not a match")
	}

	return t, nil
}
