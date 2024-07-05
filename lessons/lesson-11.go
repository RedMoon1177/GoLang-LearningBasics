// Go (Golang) Tutorial #14 - Pointers

package lessons

import "fmt"

func updateTheName(x *string) {
	*x = "wedge"
}

func Lesson11() {
	name := "tifa"

	// updateTheName(name)

	// fmt.Println("memory address of name is: ", &name)

	m := &name
	// fmt.Println("memory address: ", m)
	// fmt.Println("value at memory address: ", *m)

	fmt.Println(name)
	updateTheName(m)

	fmt.Println(name)
}

/*
|--name---|----m-----|
|  0x001  |  0x002   |
|---------|----------|
|  "tifa" |  p0x001  |
|---------|----------|
*/
