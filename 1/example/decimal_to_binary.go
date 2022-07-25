package main

import (
	"fmt"
	"test-bounties-1/app"
)

func main() {
	decimal := 101
	binary, _ := app.DecToBin(decimal)
	fmt.Printf("%d is binary form of %d", binary, decimal)
}
