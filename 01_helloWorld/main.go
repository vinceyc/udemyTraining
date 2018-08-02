package main

import "fmt"

func main() {
	fmt.Println(42)

	// Print with formatting

	// Format verbs:
	// %d means decimal
	// %b means binary
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
