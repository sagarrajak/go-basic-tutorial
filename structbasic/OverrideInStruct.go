package structbasic

import "fmt"

type Base struct{}

func (Base) Magic() { fmt.Print("base magic ") }

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() { fmt.Println("voodoo magic") }

func OverrideInStruct() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
