package suchen

import "fmt"

func ExampleContainsSum_nonempty() {
	l1 := []int{1, 14, 14, 14, 2, 23, 5}
	l2 := []int{1, 3, 5, 14, 2, 23, 17}
	l3 := []int{1, 7, 2, 3, 5, 8, 1, 15, 2}

	fmt.Println(ContainsSum(l1))
	fmt.Println(ContainsSum(l2))
	fmt.Println(ContainsSum(l3))

	// Output:
	// true
	// true
	// false
}

func ExampleContainsSum_short() {
	l1 := []int{2, 3}
	l2 := []int{14, 14, 14}

	fmt.Println(ContainsSum(l1))
	fmt.Println(ContainsSum(l2))

	// Output:
	// false
	// true
}

func ExampleContainsSum_empty() {
	l1 := []int{}

	fmt.Println(ContainsSum(l1))

	// Output:
	// false
}
