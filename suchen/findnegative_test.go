package suchen

import "fmt"

func ExampleFindNegative_nonempty() {
	l1 := []int{1, 3, -1, 7, 9}
	l2 := []int{1, 3, -1, 7, -1}
	l3 := []int{1, 3, 5, 7, -1}
	l4 := []int{1, 3, 5, 7, 9}

	fmt.Println(FindNegative(l1))
	fmt.Println(FindNegative(l2))
	fmt.Println(FindNegative(l3))
	fmt.Println(FindNegative(l4))

	// Output:
	// 2
	// 2
	// 4
	// -1
}

func ExampleFindNegative_empty() {
	l1 := []int{}

	fmt.Println(FindNegative(l1))

	// Output:
	// -1
}
