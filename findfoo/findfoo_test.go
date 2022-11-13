package findfoo

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

func ExampleFind_nonempty() {
	l1 := []int{1, 3, 5, 7, 9}

	fmt.Println(Find(l1, 3))
	fmt.Println(Find(l1, 5))
	fmt.Println(Find(l1, 7))
	fmt.Println(Find(l1, 4))
	fmt.Println(Find(l1, 6))

	// Output:
	// 1
	// 2
	// 3
	// -1
	// -1
}

func ExampleFind_empty() {
	l1 := []int{}

	fmt.Println(Find(l1, 3))
	fmt.Println(Find(l1, 5))

	// Output:
	// -1
	// -1
}

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
