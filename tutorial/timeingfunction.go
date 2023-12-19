package tutorial

func Calculation() {
	var sum float64 = 0
	for i := 0; i < 10e9; i++ {
		sum += float64(i)
	}
}
