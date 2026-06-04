package main

import "fmt"

func CalculateAge(birthYear int, currentYear int) int {
	age := currentYear - birthYear
	return age
}

func main() {
	answer := CalculateAge(1994, 2026)
	fmt.Printf("You are %d years old!", answer)
}
