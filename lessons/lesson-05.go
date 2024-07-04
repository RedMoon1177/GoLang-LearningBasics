package lessons

import "fmt"

// Go (Golang) Tutorial #7 - Loops
func Lesson05() {

	// ~ "traditional" WHILE
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	// ~ "traditional" FOR
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is: ", i)
	// }

	names := []string{"anne", "mario", "linh", "brown", "luigi"}

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(names[i])
	// }

	// ~ FOREACH
	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// if dont want to use index, replace by "_"
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)
}
