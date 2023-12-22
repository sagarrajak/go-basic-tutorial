package structbasic

import "fmt"

type T struct {
	a int
}

func (t T) printValue(message string) {
	fmt.Println(t.a)
}

func callMethod(t T, method func(T, string)) {
	method(t, "A message")
}

func MainCaller() {
	t := T{10}

	f := T.printValue

	callMethod(t, f)
}
