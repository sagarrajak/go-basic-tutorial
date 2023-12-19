package slice

import "fmt"

func sum(values []int) int {
	s := 0
	for i := 0; i < len(values); i++ {
		s += values[i]
	}
	return s
}

func SliceExample() {
	var arr [5]int
	var slice1 []int = arr[0:5]

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d\n", slice1[i])
	}

	var x = []int{2, 3, 5, 7, 11}
	t := x[:]

	slice2 := x[1:2]

	fmt.Printf("The length of array is %d\n", len(t))
	fmt.Printf("The length capacity of the array %d\n", cap(t))

	fmt.Printf("The length of array is %d\n", len(slice2))
	fmt.Printf("The capacity of array is %d\n", cap(slice2))

	fmt.Println(sum(x[:]))
}
