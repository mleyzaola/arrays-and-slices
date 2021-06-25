package main

import "fmt"

var words = []string{
	"one",
	"two",
	"three",
}

func main() {
	// slice declaration

	// declare a slice of items -capacity starts at zero, but can increase/decrease as needed-
	var items []string

	// we can append items to a slice using `append`
	items = append(items, "zero")
	fmt.Println("[zero] items with one element: ", items)
	// note the three dot notation
	items = append(items, words...)
	fmt.Println("[zero one two three] appending more: ", items)
	fmt.Println()

	// we can extract individual items from a slice, like a regular array in any other language
	//fmt.Println("[two] slice elements can be referenced by index: ", words[1])
	//fmt.Println()

	// consider the second argument minus one
	//fmt.Println("[two] also by range: ", words[1:2])
	//fmt.Println()

	// consider the second argument minus one
	//fmt.Println("[one two] another example: ", words[0:2])
	//fmt.Println()

	// we can ignore the first argument, it will be counted as zero
	//fmt.Println("[one two] another example: ", words[:2])
	//fmt.Println()

	// we can ignore the second argument, it will be counted as the last item
	//fmt.Println("[three] another example: ", words[2:])
	//fmt.Println()

	// this won't compile, first argument cannot be greater than second
	//fmt.Println("this won't work: ", words[2:1])

	// this will compile, but will panic
	//fmt.Println("i will panic and crash", words[9])
	//fmt.Println("i will panic and crash", words[2:5])

	// there is no command for item removal, but we can assign a section of a slice to it
	//items = items[1:3]
	//fmt.Println("[one two] removal of slice elements: ", items)
	//fmt.Println()

	// copying values from slices
	//var sliceCopy []string
	//copiedCount := copy(sliceCopy, words)

	// this didn't work because `copy` will take the slice with less elements, zero in this case
	//fmt.Printf("[] elements copied: %d, we copied nothing: %v\n", copiedCount, sliceCopy)
	//fmt.Println()

	// we can also copy a slice, but we need to give it a size
	// because copy will consider the minimum of the two parameters
	// so we initialize the slice with a fixed capacity
	//sliceCopy = make([]string, len(words))
	//copiedCount = copy(sliceCopy, words)
	//fmt.Printf("we copied: %d elements\n", copiedCount)
	//fmt.Println("[one two three] these elements were copied from another slice: ", words)
	//fmt.Println()

	// len() can be used to get the length of a slice
	//fmt.Println("size of words slice: ", len(words))
	//fmt.Println()

	// we can pass slices as arguments to functions
	//SliceReceiverByValue(sliceCopy)
	//fmt.Println()

	// and also we can pass a pointer of the slice
	//SliceReceiverByRef(&sliceCopy)
	//fmt.Println()

	// looping over a slice can be done using its index explicitly
	//for i := 0; i < len(words); i++ {
	//	fmt.Println("element: ", i, words[i])
	//}
	//fmt.Println()

	// but it is easier to use automatic references with a range
	// the first element of the range is the slice index
	// so, we can reference each item using the index
	//for i := range words {
	//	fmt.Println("element: ", i, words[i])
	//}
	//fmt.Println()

	// the second argument of using a range in slices, is the value of each item
	// which makes it easy to work with index and values with less code
	//for i, v := range words {
	//	fmt.Println("element: ", i, v)
	//}
	//fmt.Println()
}

func SliceReceiverByValue(v []string) {
	fmt.Printf("i received a slice with: %d elements: %v\n", len(v), v)
}

func SliceReceiverByRef(v *[]string) {
	// notice we need to de-reference the slice in order to work with it
	fmt.Printf("i received a slice with: %d elements: %v\n", len(*v), v)
}
