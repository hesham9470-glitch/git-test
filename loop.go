package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		fmt.Printf("the rocket will launch in %d!\n", i)
	}
	fmt.Printf("Blastoff!")
}
