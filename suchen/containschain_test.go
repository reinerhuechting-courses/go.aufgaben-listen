package suchen

import "fmt"

func ExampleContainsChain_nonempty() {
	l1 := []int{1, 4, 5, 6, 7, 9, 5}
	l2 := []int{1, 7, 9, 5, 4, 5, 6}
	l3 := []int{1, 3, 5, 2, 7, 9, 5}

	fmt.Println(ContainsChain(l1))
	fmt.Println(ContainsChain(l2))
	fmt.Println(ContainsChain(l3))

	// Output:
	// true
	// true
	// false
}

func ExampleContainsChain_short() {
	l1 := []int{2, 3}
	l2 := []int{1, 2, 3}

	fmt.Println(ContainsChain(l1))
	fmt.Println(ContainsChain(l2))

	// Output:
	// false
	// true
}

func ExampleContainsChain_empty() {
	l1 := []int{}

	fmt.Println(ContainsChain(l1))

	// Output:
	// false
}
