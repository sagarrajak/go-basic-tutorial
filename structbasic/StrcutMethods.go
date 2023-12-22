package structbasic

type TwoInts struct {
	A, B int
}

func (pt *TwoInts) Sum() int {
	return pt.A + pt.B
}

func (pt *TwoInts) Mul() int {
	return pt.A * pt.B
}
