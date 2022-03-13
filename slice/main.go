package main

import "fmt"

func main() {
	var s []int
	// nill => zero value
	if s == nil {
		fmt.Println("it's nil.")
	}

	a := [...]int{1, 3, 5, 7, 9}
	s = a[:]
	fmt.Printf("%d %d %p %p \n", len(s), cap(s), s, &a)
	// slice s -> point to array a

	s = append(s, 11, 13)
	fmt.Printf("%d %d %p %p \n", len(s), cap(s), s, &a)
}
