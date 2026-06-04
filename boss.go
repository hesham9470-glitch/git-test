package main

import "fmt"

func PrintUpper(word string) {
	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			fmt.Printf("%c", word[i]-32)
		} else {
			fmt.Printf("%c", word[i])
		}
	}
}

func main() {
	PrintUpper("hEsHam 2026!")
}
