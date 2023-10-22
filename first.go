package main

import (
	"fmt"
	"time"
)

var num int = 10

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
	fmt.Println("Hello world")
	var num float32 = 5.2
	const val = 5.134
	const (
		UNKNOWN = iota
		MALE    = iota
		FEMALE  = iota
	)
	var testture bool = false

	fmt.Println(num)
	fmt.Println(int(num))
	fmt.Println(val)
	fmt.Println(UNKNOWN)
	fmt.Println(MALE)
	fmt.Println(FEMALE)
	fmt.Println(testture)
	num = 12
	fmt.Println(num)
	test()
	pointers()
}
