package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		count, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word] = count + 1
		}
	}
	return m
}

func main() {
	var m map[string]int // map is 'nil'
	/* 	m["try"] = 1         // panic */

	m = WordCount("I do not do my homework, I play video games instead.")
	fmt.Println(m)

	for key, val := range m {
		fmt.Printf("key: %v, val: %v\n", key, val)
	}

	m = map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println(m)

	v, ok := m["A"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "A")
	fmt.Println(m)

	v, ok = m["A"]
	fmt.Println("The value:", v, "Present?", ok)
}
