package structbasic

import "fmt"

type inner struct {
	in1 int
	in2 int
}

type Outer struct {
	b int
	c float64
	inner
	float64
}

func HelperOuterInner() {
	outer := new(Outer)
	outer.b = 10
	outer.c = 12
	outer.in1 = 12
	outer.in2 = 13
	outer.float64 = 3234

	person := struct {
		name, surname string
	}{"barak", "obama"}

	fmt.Println(person)
	fmt.Println(outer)
}
