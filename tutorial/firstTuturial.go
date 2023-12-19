package tutorial

type flt func(int) bool

func IsOdd(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func IsEvent(n int) bool {
	if n%2 == 1 {
		return true
	}
	return false
}

func Filters(s1 []int, f flt) []int {
	var res []int
	for _, val := range s1 {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}
