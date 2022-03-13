package main

import (
	"fmt"
	"strings"
)

// TODO: split words and count them
func wordCount(s string) map[string]int {
	split := strings.Split(s, " ")
	fmt.Println(split)
	result := map[string]int{}
	for i := 0; i < len(split); i++ {
		result[split[i]] = result[split[i]] + 1
	}
	return result
}

func main() {
	// need make
	var m map[string]string
	m = make(map[string]string)

	// no make
	x := map[string]string{}

	if m == nil {
		fmt.Println("it's nil.")
	}

	m["a"] = "apple"
	m["b"] = "banana"

	x["a"] = "ant"

	s := m["a"]

	z := x["a"]
	fmt.Println(s, z)

	fmt.Print(wordCount("Apple Banana Apple Banana apple"))
}
