package slice

import "fmt"

func ReSlicing() {
	slice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice = slice[0 : i+1]
		fmt.Printf("Size of slice is %d\n", len(slice))
	}
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := arr[0:4]
	fmt.Printf("len before %d\n", len(slice2))
	slice2 = arr[2:4]
	fmt.Printf("len after %d\n", len(slice2))
}
