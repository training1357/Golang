package main

import (
	"fmt"

	str "strings" // package alias

	"example.com/packer/numbers"
	"example.com/packer/strings"
	"example.com/packer/strings/greetings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greetings.WelcomeText)
	fmt.Println(strings.Reverse("Hari Budisantoso"))
	fmt.Println(str.Count("Go emang luar biasa, Go Go Go", "Go"))
}
