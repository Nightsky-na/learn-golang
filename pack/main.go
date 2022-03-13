package main

import (
	"fmt"

	"github.com/Nightsky-na/pack/book"
)

func main() {
	b := book.New()
	// b := book.book{Name: "test"} => ไม่ได้ตัวเล็ก

	fmt.Printf("%T %v\n", b, b)

	b.Name = "Rign"
	fmt.Printf("%T %v\n", b, b)
}
