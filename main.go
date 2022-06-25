package main

import (
	"fmt"

	"github.com/husnulnawafil/cmlabs-task/palindrome"
)

func main() {
	// fizzbuzz.FizzBuzz(100)

	words := []string{"SAIPPUAKIVIKAUPPIAS", "Aibohphobia", "Anna", "Civic", "My gym", "No lemon, no melon"}
	for _, word := range words {
		result := palindrome.IsPalindrome(word)
		fmt.Println(result)
	}
}
