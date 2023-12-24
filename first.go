package main

import (
	"fmt"
	"sagarrajak/packer/structbasic"
	"strconv"
	"time"
)

var num int = 10

func forLoopTest() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print(i)
		}
	}
	// complex for loop
	for i, j := 0, 100; i < j; i, j = i+1, j-1 {
		fmt.Printf("%d, %d\n", i, j)
	}

	j := 10

	for j < 100 {
		fmt.Println()
		j = j + 1
	}
}

func switchCase() {
	test := 12
	// switch test {
	// case 15:
	// 	fmt.Println("This is something else")
	// case 12, 13:
	// 	fmt.Println("This is another swithc statement!")
	// default:
	// 	fmt.Println("None of this condition matched!")
	// }
	switch {
	case test > 10 && test < 13:
		fmt.Println("Test is between 10 and 13")
	case test < 9:
		fmt.Println("Test us below nine")
	default:
		fmt.Println("This is default")
	}

	// init with switch statment
	switch test2 := 100; {
	case test2 < 120:
		fmt.Println("This is below 120")
	default:
		fmt.Println("Something default is having printed here!")
	}
}

func errorFunction() {
	var org string = "TEST"
	var an int
	var err error
	an, err = strconv.Atoi(org)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(an)
}

func controlStatment() {
	n := 43
	if n%2 > 0 {
		fmt.Println("this is statement")
	} else {
		fmt.Println("this is false")
	}

	if f := 25; f > 10 {
		fmt.Println("new value is grether then 10")
	}
}

func pointers() {
	var i2 = 5
	var address = &i2
	fmt.Printf("value %d and address is %p\n", i2, &i2)
	fmt.Printf("%p address and value %d\n", address, *address)
	var str = "String"
	var straddress *string = &str
	fmt.Printf("previous value is \"%s\"\n", str)
	*straddress = "new Value"
	fmt.Printf("new value is \"%s\"\n", str)

}

func test() {
	t := time.Now()
	t_utc := time.Now().UTC()
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	fmt.Printf("%02d.%02d.%4d\n", t_utc.Day(), t_utc.Month(), t_utc.Year())
	fmt.Println(t_utc)
	s := t.Format("2006 01 02")
	println(s)
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	fmt.Println(t.Format(time.ANSIC))
}

func main() {
	// fmt.Println("Hello world")
	// var num float32 = 5.2
	// const val = 5.134
	// const (
	// 	UNKNOWN = iota
	// 	MALE    = iota
	// 	FEMALE  = iota
	// )
	// var testture bool = false

	// fmt.Println(num)
	// fmt.Println(int(num))
	// fmt.Println(val)
	// fmt.Println(UNKNOWN)
	// fmt.Println(MALE)
	// fmt.Println(FEMALE)
	// fmt.Println(testture)
	// num = 12
	// fmt.Println(num)
	// test()
	// pointers()
	// controlStatment()
	// errorFunction()
	// switchCase()
	// forLoopTest()
	// slice := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(tutorial.Filters(slice, tutorial.IsEvent))
	// fmt.Println(tutorial.Filters(slice, tutorial.IsOdd))
	// fmt.Println(tutorial.TestAny())
	// callFunc := tutorial.GetInc(10)
	// getFib := tutorial.GetFib()
	// fmt.Println(callFunc(12))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(getFib())
	// }
	// builder := tutorial.StringBulder()
	// builder("temp")
	// builder("temp2")
	// build := builder("temp3")
	// fmt.Println(build())

	// // timeing calculation
	// start := time.Now()
	// tutorial.Calculation()
	// end := time.Now()
	// delta := end.Sub(start)
	// fmt.Printf("Time change from this is\n", delta)

	// array.BasicOperation()
	// array.OperationsExamples()
	// slice.SliceExample()
	// slice.SliceWithMake()
	// slice.TwoDSlices()
	// slice.RangeConstruct()
	// slice.ReSlicing()
	// slice.CopyInGO()
	// maptut.MapBasic()

	// // node.right = &structbasic.NewNode(14, nil, nil)
	// node := structbasic.NewNode(65, nil, nil)
	// node.Left = structbasic.NewNode(12, nil, nil)
	// node.Right = structbasic.NewNode(12, nil, nil)
	// fmt.Printf("module data node %f\n", node.Data)
	// fmt.Printf("module data node %f\n", node.Left.Data)
	// fmt.Printf("module data node %f\n", node.Right.Data)

	// structbasic.HelperOuterInner()
	// twoPoint := &structbasic.TwoInts{A: 12, B: 12}
	// fmt.Printf("Sum of two ponter %d\n", twoPoint.Sum())
	// fmt.Printf("Mul of two pointer %d\n", twoPoint.Mul())
	// structbasic.TestMethodAsPointer()
	// structbasic.MainCaller()
	// person := new(structbasic.Person)
	// person.SetFirstName("sagar")
	// fmt.Println(person.FirstName())
	// snakegame.SnakeGame()
	// tp := new(structbasic.TwoPointer)
	// tp.SetFirst(12)
	// tp.SetSecond(20)
	// fmt.Println(tp)
	// point := &structbasic.NamedPoint{Point: structbasic.Point{X: 2.2, Y: 4.4}, Name: "sagar"}
	// fmt.Println(point.Abs())
	// structbasic.CompositionStruct()
	structbasic.OverrideInStruct()
}
