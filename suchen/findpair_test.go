package suchen

import "fmt"

func ExampleFindPair_nonempty() {
	l1 := []int{1, 3, 3, 7, 9, 1, 5, 5, 2}

	fmt.Println(FindPair(l1, 3))
	fmt.Println(FindPair(l1, 5))
	fmt.Println(FindPair(l1, 7))
	fmt.Println(FindPair(l1, 4))
	fmt.Println(FindPair(l1, 6))

	// Output:
	// 1
	// 6
	// -1
	// -1
	// -1
}

func ExampleFindPair_empty() {
	l1 := []int{}

	fmt.Println(FindPair(l1, 3))
	fmt.Println(FindPair(l1, 5))

	// Output:
	// -1
	// -1
}
