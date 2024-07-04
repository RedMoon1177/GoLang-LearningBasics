package lessons

import (
	"fmt"
	"math"
)

// Go (Golang) Tutorial #9 - Using Functions

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func Lesson06() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
