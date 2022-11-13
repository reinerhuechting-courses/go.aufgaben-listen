package suchen

import "fmt"

func ExampleContainsMultiple_nonempty() {
	l1 := []int{1, 3, 5, 3, 7, 9, 5}

	fmt.Println(ContainsMultiple(l1, 3))
	fmt.Println(ContainsMultiple(l1, 5))
	fmt.Println(ContainsMultiple(l1, 7))
	fmt.Println(ContainsMultiple(l1, 4))
	fmt.Println(ContainsMultiple(l1, 6))

	// Output:
	// true
	// true
	// false
	// false
	// false
}

func ExampleContainsMultiple_empty() {
	l1 := []int{}

	fmt.Println(ContainsMultiple(l1, 3))
	fmt.Println(ContainsMultiple(l1, 5))

	// Output:
	// false
	// false
}
