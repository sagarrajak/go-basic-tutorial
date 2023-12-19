package maptut

import "fmt"

func SliceOfMap() {
	items := make([]map[int]int, 100)
	for _, item := range items {
		item = make(map[int]int, 1)
		item[1] = 1000
	}

	for _, item := range items {
		fmt.Printf("%d\n", item[1])
	}
}
