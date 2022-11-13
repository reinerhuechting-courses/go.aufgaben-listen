package suchen

import "fmt"

func ExampleContains_nonempty() {
	l1 := []int{1, 3, 5, 7, 9}

	fmt.Println(Contains(l1, 3))
	fmt.Println(Contains(l1, 5))
	fmt.Println(Contains(l1, 7))
	fmt.Println(Contains(l1, 4))
	fmt.Println(Contains(l1, 6))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleContains_empty() {
	l1 := []int{}

	fmt.Println(Contains(l1, 3))
	fmt.Println(Contains(l1, 5))

	// Output:
	// false
	// false
}
