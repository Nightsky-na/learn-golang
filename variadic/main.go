package main

import "fmt"

func main() {
	variadicString("a", "b")
	s := []string{"a", "b", "c"}
	//spread operator
	variadicString(s...)
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
