package slice

import (
	"fmt"
	"strings"
)

func RangeConstruct() {
	slice1 := make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 20
	}

	for idx, value := range slice1 {
		fmt.Printf("indx is %d and value is %d\n", idx, value)
	}

	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for idx, season := range seasons {
		seasons[idx] = strings.ToUpper(season)
		fmt.Printf("%s\n", seasons[idx])
	}
	// range with multi -  dimencion

	screen := [10][10]int{}

	count := 0

	for i, row := range screen {
		for j, _ := range row {
			screen[i][j] = count
			count++
		}
	}

	for _, row := range screen {
		fmt.Println(row)
	}
}
