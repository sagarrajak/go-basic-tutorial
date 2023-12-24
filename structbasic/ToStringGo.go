package structbasic

import "strconv"

type TwoPointer struct {
	first  int
	second int
}

func (tp *TwoPointer) String() string {
	return "{" + "first:" + strconv.Itoa(tp.first) + ",second:" + strconv.Itoa(tp.second) + "}"
}

func (tp *TwoPointer) SetFirst(val int) {
	tp.first = val
}

func (tp *TwoPointer) SetSecond(val int) {
	tp.second = val
}
