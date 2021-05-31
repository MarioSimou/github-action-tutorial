package main

import (
	"fmt"
	"github.com/MarioSimou/calculator/internal"
)

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	var sum = internal.Sum(nums...)

	fmt.Printf("Total: %d\n", sum)
}
