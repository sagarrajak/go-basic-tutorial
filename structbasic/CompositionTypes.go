package structbasic

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Log
	name string
}

func (lg *Log) Add(mes string) {
	lg.msg += "\n" + mes
}

func (lg *Log) String() string {
	return lg.msg
}

// func (cus *Customer) Log() *Log {
// 	return cus.log
// }

func CompositionStruct() {
	cus := new(Customer)
	cus.Add("some message")
	cus.Add("another message")
	fmt.Println(cus.msg)
}
