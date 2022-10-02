package main

import "fmt"

func intro() {
	bools := []bool{true, false, true, true, false, true}
	printSlice("bools", bools)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("s", s)

	printSlice("s[1:5]", s[1:5])
	printSlice("s[:3]", s[:3])
	printSlice("s[1:]", s[1:])
}

func references() {
	fmt.Println("\n# Slices referencing other slices")

	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	printSlice("names", names)

	fmt.Println("a and b are subsets of 'names'")
	a := names[0:2]
	b := names[1:3]
	printSlice("a", a)
	printSlice("b", b)

	fmt.Println("Modifying b[0]")
	b[0] = "XXX"
	printSlice("a", a)
	printSlice("b", b)
	printSlice("names", names)
}

func using_make() {
	fmt.Println("\n# Creating slices with make")

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
}

func two_dimensional() {
	fmt.Println("\n# Slice of slices")
	board := [][]string{
		{"a", "b", "c"},
		{"x", "y", "z"},
		{"*", "!", "$"},
	}
	// range keyword iterates over a slice
	for i, line := range board {
		fmt.Printf("line %d: %v\n", i, line)
	}
}

func appending() {
	fmt.Println("\n# Appending to slices")
	var s []int
	printSlice("s", s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice("s", s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice("s", s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice("s", s)
}

func main() {
	intro()
	references()
	using_make()
	two_dimensional()
	appending()
}

func printSlice[T any](name string, s []T) {
	fmt.Printf("%s: len=%d cap=%d %v\n", name, len(s), cap(s), s)
}
