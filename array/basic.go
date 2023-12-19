package array

import "fmt"

func passByValue(arr [5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Print by pass by value %d\n", arr[i])
	}
}

func passByRef(arr *[5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Print by pass by ref %d\n", arr[i])
	}
}

func BasicOperation() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i * 2
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d th value is %d\n", i, arr[i])
	}

	var pointerArrr = new([5]int)
	for i := 0; i < 5; i++ {
		pointerArrr[i] = 10
	}
	for i := 0; i < 5; i++ {
		fmt.Println(pointerArrr[i])
	}

	passByRef(pointerArrr)
	passByRef(&arr)

	passByValue(arr)
	passByValue(*pointerArrr)
}
