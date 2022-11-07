package main

import (
	"fmt"
)

func variable_declaration() {
	// Declaring Variable
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)

	// Multiple Declarations
	var (
		employeeId          int    = 5
		firstName, lastName string = "Uzumaki", "Naruto"
	)
	fmt.Println(employeeId, firstName, lastName)

	// Short variable declaration
	name := "Bill Gates"
	age, salary, isProgrammer := 35, 50000.0, true
	fmt.Println(name, age, salary, isProgrammer)
}

func type_inference() {
	// Type Inference
	var name = "Steve Jobs"
	fmt.Printf("Variable 'name' is of type %T\n", name)

	// Multiple variable declaration with inference types
	var firstName, lastName, age, salary = "James", "Bond", 28, 75000.0
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T", firstName, lastName, age, salary)
}

func constant_declaration() {
	// Untyped Constant
	const myFavLanguage = "Kotlin"
	const sunRisesInTheEast = true

	// Multiple declaration
	const country, code = "Indonesia", 62

	const (
		employeeId string  = "E101"
		salary     float64 = 50000.0
	)

	// Typed Constant
	const typedInt int = 1234
	const typedStr string = "Hi"

	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeId, salary, typedInt, typedStr)
}

func ioata_declaration() {
	const (
		first = iota
		second
	)

	const (
		third = iota
		fourth
	)

	fmt.Println(first, second, third, fourth)
}

func pointer_declaration() {
	var a = 5.67
	var p = &a

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p = ", p)
}

func main() {
	fmt.Println("Primitive Data Types")
	// variable_declaration()
	// type_inference()
	// constant_declaration()
	//ioata_declaration()
	pointer_declaration()
}
