package slice

import "fmt"

func SliceWithMake() {

	// slice func with make function
	slice3 := make([]int, 10, 100)
	for i := 0; i < len(slice3); i++ {
		slice3[i] = i
	}

	for i := 0; i < len(slice3); i++ {
		fmt.Printf("Element at %d\n", slice3[i])
	}

	fmt.Printf("Slice func with make length %d\n", len(slice3))
	fmt.Printf("Slice func with make capacity %d\n", cap(slice3))

	// using new function
	slice4 := new([100]int)[10:20]

	for i := 0; i < len(slice4); i++ {
		slice4[i] = i*10 + i
	}

	for i := 0; i < len(slice4); i++ {
		fmt.Printf("This is simple slice with new %d\n", slice4[i])
	}
}
