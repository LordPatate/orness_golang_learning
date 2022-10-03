package main

import "fmt"

func main() {
	const n = 99999

	count := 0
	for i := 0; i < n; i++ {
		func() { count++ }()
	}
	fmt.Println(count)

	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func() {
			count--
			done <- true
		}()
	}
	for i := 0; i < n; i++ {
		<-done
	}
	fmt.Println(count)
}
