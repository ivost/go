package main

import (
	"fmt"
	"strings"
)

func ExampleRepeat() {
	fmt.Println(strings.Repeat("a", 3))
	// Output: aaa
}

func ExampleFind1() {
	a := [] int {1, 2, 3}
	fmt.Println(find(a, 2))
	// Output: 1
}

func ExampleFind2() {
	a := [] int {1, 2, 3}
	fmt.Println(find(a, 1))
	// Output: 0
}

func ExampleFind3() {
	a := [] int {1, 2, 3}
	fmt.Println(find(a, 3))
	// Output: 2
}

func ExampleFind4() {
	a := [] int {1, 2, 3}
	fmt.Println(find(a, 5))
	// Output: -1
}

func ExampleFind5() {
	a := [] int {10, 12, 15, 20, 30, 40}
	fmt.Println(find(a, 30))
	// Output: 4
}

func ExampleFind6() {
	a := [] int {10, 12, 15, 20, 30, 40}
	fmt.Println(find(a, 10))
	// Output: 0
}

func ExampleFind7() {
	a := [] int {10, 12, 15, 20, 30, 40}
	fmt.Println(find(a, 1))
	// Output: -1
}

func ExampleFind8() {
	a := [] int {10, 12, 15, 20, 30, 40}
	fmt.Println(find(a, 25))
	// Output: -1
}
