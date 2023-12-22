package structbasic

type Person struct {
	firstName string
	lastName  string
}

func (p Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(val string) {
	p.firstName = val
}
