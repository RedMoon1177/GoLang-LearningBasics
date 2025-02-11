// Go (Golang) Tutorial #13 - Pass By Value

// Go is considered a "pass by value" language because every function call receives a copy of the original value,
// whether it is a simple data type, a struct, or an array.
// This ensures that modifications inside the function do not affect the original variable, except when dealing with slices, maps, and pointers,
// where the underlying data structures or memory locations can be modified.
// This model simplifies reasoning about data flow and state changes in Go programs.
// for example:
// func modifyValue(val int) {
//     val = 10
// }

// func main() {
//     x := 5
//     modifyValue(x)
//     fmt.Println(x) // Output: 5, because x was passed by value
// }

package lessons

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func Lesson10() {

	// group A types -> stirngs, ints, bools, floats, arrays, structs (called - "Non-Pointer Values")
	name := "tifa"

	updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions (called - "Pointer Wrapper Values")
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
