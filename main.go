package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	Name     string
	Birthday time.Time
}

func (p *Person) String() string {
	return "[" +
		p.Birthday.Format("03-02-2006") +
		"] " +
		p.Name
}

var persons = []Person{
	{
		Name:     "Rebeca",
		Birthday: time.Date(1995, 8, 16, 0, 0, 0, 0, time.UTC),
	},
	{
		Name:     "Mark",
		Birthday: time.Date(1968, 9, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		Name:     "Tony",
		Birthday: time.Date(2002, 3, 21, 0, 0, 0, 0, time.UTC),
	},
	{
		Name:     "Susan",
		Birthday: time.Date(1999, 12, 5, 0, 0, 0, 0, time.UTC),
	},
}

func main() {
	// sorting primitive integers and strings is straightforward
	var numbers = []int{3, 23, 1, 19, 12}
	sort.Ints(numbers)
	fmt.Println(numbers)

	var words = []string{"the", "fox", "and", "the", "grapes"}
	sort.Strings(words)
	fmt.Println(words)

	// sorting structs requires more work
	// another way is using the sort.Slice() function
	PrintPersons(persons)
	fmt.Println()
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Birthday.Before(persons[j].Birthday)
	})
	PrintPersons(persons)

	// one way of sorting slices is implementing the sort.Sort interface
	// this requires to implement a custom type as well
	//var sortedPersons = SortedPerson(persons)
	//sort.Sort(sortedPersons)
}

func PrintPersons(persons []Person) {
	for _, v := range persons {
		fmt.Println(v.String())
	}
}

// SortedPerson is a custom type for implementing sort on Person slice
// TODO Exercise: implement sort.Sort interface
//type SortedPerson []Person
