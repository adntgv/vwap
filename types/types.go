package types

type Transaction struct {
	ProductID string
	Size      float64
	Price     float64
}

type VWAP struct {
	ProductID string
	Value     float64
}

type BufferedType struct {
	SizePrice float64
	Size      float64
}
