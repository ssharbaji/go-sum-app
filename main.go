package main

import (
	"flag"
	"fmt"
)

func Sum(a, b int) int {
	return a + b
}

func main() {
	var a, b int
	flag.IntVar(&a, "a", 0, "First integer")
	flag.IntVar(&b, "b", 0, "Second integer")

	flag.Parse()

	result := Sum(a, b)

	fmt.Printf("Sum(%v, %v) = %v", a, b, result)
}
