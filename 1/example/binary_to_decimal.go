package main

import (
	"fmt"
	"test-bounties-1/app"
)

func main() {
	binary := 110011
	decimal, _ := app.BinToDec(binary)
	fmt.Printf("%d is decimal form of %d", decimal, binary)
}
