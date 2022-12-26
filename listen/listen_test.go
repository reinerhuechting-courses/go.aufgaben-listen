package listen

import "fmt"

func ExampleCountElement() {
	l1 := []int{13, 2, 4, 25, 4, 12, 2, -3, 15, 25, 13, 4, 75, 42, 2, 4, 13, 10, 4, 17, 38}

	fmt.Println(CountElement(l1, 13))
	fmt.Println(CountElement(l1, 2))
	fmt.Println(CountElement(l1, 4))
	fmt.Println(CountElement(l1, 25))
	fmt.Println(CountElement(l1, 38))
	fmt.Println(CountElement(l1, 7))

	// Output:
	// 3
	// 3
	// 5
	// 2
	// 1
	// 0
}

func ExampleCountGreaterThan() {
	l1 := []int{13, 2, 4, 25, 4, 12, 2, -3, 15, 25, 13, 4, 75, 42, 2, 4, 13, 10, 4, 17, 38}

	fmt.Println(CountGreaterThan(l1, 13))
	fmt.Println(CountGreaterThan(l1, 2))
	fmt.Println(CountGreaterThan(l1, 4))
	fmt.Println(CountGreaterThan(l1, 25))
	fmt.Println(CountGreaterThan(l1, 38))
	fmt.Println(CountGreaterThan(l1, 7))

	// Output:
	// 7
	// 17
	// 12
	// 3
	// 2
	// 12
}

func ExampleCountNotN() {
	l1 := []int{13, 2, 4, 25, 4, 12, 2, -3, 15, 25, 13, 4, 75, 42, 2, 4, 13, 10, 4, 17, 38}

	fmt.Println(CountNotN(l1, 13))
	fmt.Println(CountNotN(l1, 2))
	fmt.Println(CountNotN(l1, 4))
	fmt.Println(CountNotN(l1, 25))
	fmt.Println(CountNotN(l1, 38))
	fmt.Println(CountNotN(l1, 7))

	// Output:
	// 18
	// 18
	// 16
	// 19
	// 20
	// 21
}

func ExampleCountDivisors() {
	l1 := []int{13, 2, 4, 25, 4, 12, 2, -3, 15, 25, 13, 4, 75, 42, 2, 4, 13, 10, 4, 17, 38}

	fmt.Println(CountDivisors(l1, 13))
	fmt.Println(CountDivisors(l1, 2))
	fmt.Println(CountDivisors(l1, 4))
	fmt.Println(CountDivisors(l1, 25))
	fmt.Println(CountDivisors(l1, 38))
	fmt.Println(CountDivisors(l1, 7))

	// Output:
	// 3
	// 12
	// 6
	// 3
	// 1
	// 1
}

func ExampleCountEqualElements_equalLength() {
	l1 := []int{1, 1, 1, 1, 1}
	l2 := []int{1, 2, 3, 1, 2}
	l3 := []int{3, 2, 3, 4, 1}

	fmt.Println(CountEqualElements(l1, l2))
	fmt.Println(CountEqualElements(l1, l3))
	fmt.Println(CountEqualElements(l2, l3))
	fmt.Println(CountEqualElements(l1, l1))

	// Output:
	// 2
	// 1
	// 2
	// 5
}

func ExampleCountEqualElements_differentLength() {
	l1 := []int{1, 1, 1, 1, 1}
	l2 := []int{1, 2, 3}

	fmt.Println(CountEqualElements(l1, l2))
	fmt.Println(CountEqualElements(l2, l1))

	// Output:
	// 1
	// 1
}

func ExampleCountGreaterElements_equalLength() {
	l1 := []int{1, 1, 1, 1, 1}
	l2 := []int{1, 2, 3, 1, 2}
	l3 := []int{3, 2, 3, 4, 1}

	fmt.Println(CountGreaterElements(l1, l2))
	fmt.Println(CountGreaterElements(l1, l3))
	fmt.Println(CountGreaterElements(l2, l3))
	fmt.Println(CountGreaterElements(l3, l2))
	fmt.Println(CountGreaterElements(l1, l1))

	// Output:
	// 0
	// 0
	// 1
	// 2
	// 0
}

func ExampleCountGreaterElements_differentLength() {
	l1 := []int{3, 4, 5}
	l2 := []int{1, 2, 6, 7, 8}

	fmt.Println(CountGreaterElements(l1, l2))
	fmt.Println(CountGreaterElements(l2, l1))

	// Output:
	// 2
	// 1
}
