package structbasic

import "fmt"

type myStruct struct {
	IntField    int
	FloatField  float64
	StringField string
}

func modifyStruct(s *myStruct) {
	s.IntField = 20
	s.StringField = "Modified"
}

func ModifyStructMain() {
	original := new(myStruct)
	original.FloatField = 10.12
	original.IntField = 12
	original.StringField = "test123"

	modifyStruct(original)

	fmt.Println("Original struct:", original)
}
