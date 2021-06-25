package main

import "fmt"

func main() {

	// maps

	// declare a map uninitialized
	var kv map[string]int

	// initialize the map
	kv = make(map[string]int)

	// assign values to the map
	kv["one"] = 1
	kv["two"] = 2

	// safe check if value exist
	if v, ok := kv["one"]; ok {
		fmt.Println("i exist: ", v)
	}

	// bad practice, can lead to unexpected results
	v := kv["three"]
	fmt.Println("because we are using a primitive type, we get its zero value", v)

	// maps can be also initialized on declaration
	var anotherMap = map[string]int{
		"one": 1,
	}
	fmt.Println("anotherMap content: ", anotherMap)
}

// SimpleDupCheckFunc will return true if there are duplicated items in `values`
// complete the function using a map
func SimpleDupCheckFunc(values []string) bool {
	kmap := make(map[string]struct{})
	for _, v := range values {
		if _, ok := kmap[v]; ok {
			return true
		}
		kmap[v] = struct{}{}
	}
	return false
}
