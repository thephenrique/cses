package main

import (
	"fmt"
	"log"
)

func main() {
	var input int

	// STDIN.
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d ", input)

	for input != 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}

		fmt.Printf("%d ", input)
	}

	fmt.Print("\n")
}
