package suchen

import "fmt"

func ExampleContainsChain2_nonempty() {
	l1 := []int{1, 5, 6, 2, 7, 9, 5}
	l2 := []int{1, 7, 3, 5, 4, 5, 6}
	l3 := []int{1, 3, 2, 5, 4, 9, 5}

	fmt.Println(ContainsChain2(l1))
	fmt.Println(ContainsChain2(l2))
	fmt.Println(ContainsChain2(l3))

	// Output:
	// true
	// true
	// false
}

func ExampleContainsChain2_short() {
	l1 := []int{2, 3}
	l2 := []int{1, 3, 5}

	fmt.Println(ContainsChain2(l1))
	fmt.Println(ContainsChain2(l2))

	// Output:
	// false
	// true
}

func ExampleContainsChain2_empty() {
	l1 := []int{}

	fmt.Println(ContainsChain2(l1))

	// Output:
	// false
}
