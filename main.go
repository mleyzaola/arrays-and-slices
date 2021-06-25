package main

type Person struct {
	Name string
}

func (p *Person) String() string {
	return p.Name
}

var personValues = []Person{
	{
		Name: "Susan",
	},
}

var personRefs = []*Person{
	{
		Name: "Rebeca",
	},
}

func main() {

	// slice management

}

func StrToPtr(v string) *string {
	return &v
}
