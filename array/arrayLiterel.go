package array

import "fmt"

func OperationsExamples() {
	var arr = [...]int{1, 2, 3, 4, 56}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\n", arr[i])
	}
	var stringArray = [...]string{1: "sagar", 2: "sumit"}
	for i := 0; i < len(stringArray); i++ {
		fmt.Printf("%d ubde %s\n", i, stringArray[i])
	}
}
