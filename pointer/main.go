package main

import "fmt"

func main() {
	d := 2
	double(&d)

	fmt.Println(d)
}

// func double(d *int) {
// 	*d = *d * 2
// }

// n := 2
// double(&n)
// n == 4
func double(n *int) {
	*n = *n * 2
}

// name := "Bob"
// appendGreeting(&name)
// name == "Hi, Bob"
func appendGreeting(s *string) {
	fmt.Sprintf("Hi, %d", *s)
}
