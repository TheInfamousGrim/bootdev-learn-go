package main

import "fmt"

func main() {
	accountAge := 2.6
	accountAgeInt := int64(accountAge)

	// Create a new "accountAgeInt" here
	// It should be the result of casting "accountAge" to an integer

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
