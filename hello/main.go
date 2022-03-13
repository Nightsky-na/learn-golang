package main

import "fmt"

// greeting("Pallat") should return "Hello, Pallat"
func greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name,age,drink)
}

func isOdd(n int) bool {
	if n % 2 == 1 {
		return true
	} else {
		return false
	}
}

// It should return sum of n, n-1, n-2, ..., 1
// sumOfFirst(3) should return 3+2+1=6
func sumOfFirst(n int) int {
	var sum_num int = 0
	for i := n; i > 0; i-- {
		sum_num += i
	}
	return sum_num
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1 // 1 is not a prime number
}

func main() {
	fmt.Println("Hello, Kontawat")
	fmt.Println(greetingWithAge("Meng",19))
	fmt.Println(sumOfFirst(5))
}