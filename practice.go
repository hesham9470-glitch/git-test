package main

import "fmt"

func main() {
	lvl := 5

	if lvl > 2 {
		fmt.Printf("Hey! your level is too high and is above level %d", lvl)
	} else {
		fmt.Printf("Hey your level is lower than %d keep grinding!", lvl)
	}
}
