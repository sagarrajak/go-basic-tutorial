package slice

import "fmt"

func CopyInGO() {
	// sl_from := []int{1, 2, 3}
	// sl_to := make([]int, 10)
	// n := copy(sl_to, sl_from)
	// fmt.Printf("number of items copied %d\n", n)
	// fmt.Println(sl_to)
	// fmt.Println()

	originalSlice := []string{"a", "b", "c", "d", "e"}
	insertionSlice := []string{"x", "y", "z"}
	index := 2

	result := insertSlice(originalSlice, insertionSlice, index)
	fmt.Println(result) // Output: [a b x y z c d e]
}

func insertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(insertion)+len(slice))
	n := copy(result, slice[0:index])
	n += copy(result[index:], insertion)
	copy(result[n:], slice[index:])
	return result
}

func Filter(s []int, fn func(int) bool) []int {
	out := []int{}
	for _, val := range s {
		if fn(val) {
			out = append(out, val)
		}
	}

	return out
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func reverse(str string) string {
	s := []byte(str)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		tem := s[i]
		s[i] = s[j]
		s[j] = tem
	}
	return string(s)
}
