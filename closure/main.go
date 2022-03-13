package main

import "fmt"

func main() {
	fn := newCounterFunc()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

// return fn that fn return int
// write same args
func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
