package slice

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func TwoDSlices() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}

	values := [][]int{}

	values = append(values, []int{1, 2, 3})
	values = append(values, []int{4, 5, 6})

	fmt.Println("First row")
	fmt.Println(values[0])
	fmt.Println("Second row")
	fmt.Println(values[1])
	fmt.Println("First row, first element!")
	fmt.Println(values[0][0])
	fmt.Println("Second row, first elemenmt1")
	fmt.Println(values[1][0])
}
