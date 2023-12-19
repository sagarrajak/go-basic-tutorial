package tutorial

func TestAny() float64 {
	Exc := func() float64 {
		var sum float64 = 0
		for i := 0; i <= 1e10; i++ {
			sum += float64(i)
		}
		return sum
	}
	return Exc()
}

func GetInc(n int) func(x int) int {
	return func(x int) int {
		return x + n
	}
}

func GetFib() func() int {
	x, y := 1, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func StringBulder() func(str string) func() string {
	buildingStr := ""
	return func(str string) func() string {
		buildingStr += str
		return func() string {
			var temp string = buildingStr
			buildingStr = ""
			return temp
		}
	}

}
