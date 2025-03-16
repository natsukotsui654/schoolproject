
// This is a Go program that generates a random number between 1 and 10
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(10) + 1
	fmt.Println("Random number:", result)
}