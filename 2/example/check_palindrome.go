package main

import (
	"fmt"
	"test-bounties-2/app"
)

func main() {
	str := "aku suka makan nasi"
	str2 := "di rumah saya ada kasur rusak"
	str3 := "abcde edcbza"

	fmt.Printf("Input '%s' output '%s'. \n", str, app.Palindrome(str))
	fmt.Printf("Input '%s' output '%s'. \n", str2, app.Palindrome(str2))
	fmt.Printf("Input '%s' output '%s'. \n", str3, app.Palindrome(str3))
}
