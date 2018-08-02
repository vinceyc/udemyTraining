package main

import "fmt"

func main() {
	for i := 60; i < 121; i++ {
		fmt.Printf("%d \t %b \t %q \n", i, i, i)
	}
}
