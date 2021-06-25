package main

type Person struct {
	Name     string
	LastName *string
}

func (p *Person) String() string {
	if p.LastName == nil {
		return p.Name
	}
	return p.Name + " " + *p.LastName
}

var personValues = []Person{
	{
		Name:     "Susan",
		LastName: StrToPtr("Smith"),
	},
}

var personRefs = []*Person{
	{
		Name:     "Rebeca",
		LastName: StrToPtr("Zan"),
	},
}

func main() {

	// slices and pointers

	// pass a value to slice of values
	//SliceByValue(personValues)
	//fmt.Println(personValues[0].String())

	// pass a value to slice of refs
	//SliceByValueOfPointers(personRefs)
	//fmt.Println(personRefs[0].String())

	//SliceLoopValues(personRefs)
	//fmt.Println(personRefs[0].String())

	// pass a slice by ref but with value on its contents
	//SliceLoopPtrReferenceSlice(&personValues)
	//fmt.Println(personValues[0].String())

	//SliceLoopPtrReferenceSlice2(&personValues)
	//for i, v := range personValues {
	//	fmt.Println(i, ":", v.String())
	//}
}

func StrToPtr(v string) *string {
	return &v
}

func SliceByValue(persons []Person) {
	susan := persons[0]
	susan.LastName = StrToPtr("I won't change")
}

func SliceByValueOfPointers(persons []*Person) {
	rebeca := persons[0]
	rebeca.LastName = StrToPtr("Changed!")
}

func SliceLoopValues(persons []*Person) {
	for _, v := range persons {
		v.Name = "Changed!"
	}
}

func SliceLoopPtrReferenceSlice(persons *[]Person) {
	for _, v := range *persons {
		v.Name = "I won't change"
	}
}

func SliceLoopPtrReferenceSlice2(persons *[]Person) {
	*persons = append(*persons, Person{
		Name:     "New",
		LastName: StrToPtr("Person"),
	})
}
