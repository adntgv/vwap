package calculator

import (
	"fmt"
	"strings"
	"testing"

	"github.com/adntgv/vwap/types"
)

type testCase struct {
	transaction  *types.Transaction
	expectedVWAP float64
}

var testTransactions = []testCase{
	{
		transaction: &types.Transaction{
			ProductID: "test",
			Size:      0.1,
			Price:     0.2,
		},
		expectedVWAP: 0.1 * 0.2 / 0.1,
	}, {
		transaction: &types.Transaction{
			ProductID: "test",
			Size:      0.3,
			Price:     0.4,
		},
		expectedVWAP: (0.1*0.2 + 0.3*0.4) / (0.1 + 0.3),
	}, {
		transaction: &types.Transaction{
			ProductID: "test",
			Size:      0.5,
			Price:     0.6,
		},
		expectedVWAP: (0.1*0.2 + 0.3*0.4 + 0.5*0.6) / (0.1 + 0.3 + 0.5),
	},
}

func TestCalculator_Calculate(test *testing.T) {
	numTest := len(testTransactions)
	calculator := newCalculator(numTest, nil, nil)

	for _, t := range testTransactions {
		vwap := calculator.process(t.transaction)

		expectedPrefix := fmt.Sprint(t.expectedVWAP)

		if strings.HasPrefix(fmt.Sprint(vwap), expectedPrefix) {
			test.Errorf("resulting vwap %v, expected %v", vwap.Value, t.expectedVWAP)
		}
	}
}
