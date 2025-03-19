package main

import "fmt"

func main() {
	x := 42
	if x > 10 {
		fmt.Println("The answer is greater than 10")
	} else if x < 10 {
		fmt.Println("The answer is less than 10")
	} else {
		fmt.Println("The answer is equal to 10")
	}
}
