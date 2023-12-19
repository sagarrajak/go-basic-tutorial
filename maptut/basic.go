package maptut

import "fmt"

func MapBasic() {
	mapList := map[string]int{"one": 1, "two": 2}
	keys := make([]string, 0, len(mapList))
	for k := range mapList {
		keys = append(keys, k)
	}

	mapFunc := map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		},
	}

	temp := mapFunc[1]()
	fmt.Println(temp)

	mapStringFloat := make(map[string]float64, 100)
	strSlic := [...]string{"one", "two", "three"}
	for _, str := range strSlic {
		mapStringFloat[str] = 23234.234423
	}

	for temp2 := range mapStringFloat {
		fmt.Println(temp2, mapStringFloat[temp2])
	}

	// check if key exist

	value, ifPresent := mapStringFloat["test123"]

	if ifPresent {
		fmt.Printf("present value is %d\n", value)
	} else {
		fmt.Printf("key %s is not present in array\n", "test123")
	}

	fmt.Println(mapStringFloat["test123"])
}
