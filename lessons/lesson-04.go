package lessons

import (
	"fmt"
	"sort"
)

// Go (Golang) Tutorial #6 - The Standard Library
func Lesson04() {
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	// fmt.Println("original string value = ", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90)
	fmt.Println(index)

	names := []string{"anne", "mario", "linh", "brown", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

}
