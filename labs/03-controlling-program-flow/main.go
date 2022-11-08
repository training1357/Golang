package main

import "fmt"

func main() {

	// if statement
	var x = 25
	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	// if with a short statement
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

	// if-else statement
	var age = 15
	if age >= 17 {
		fmt.Println("Anda boleh ikut pemilu")
	} else {
		fmt.Println("Anda belum boleh ikut pemilu")
	}

	// switch statement
	var dayOfWeek = 6

	switch dayOfWeek {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
		fmt.Println("Weekend")
	case 7:
		fmt.Println("Minggu")
		fmt.Println("Weekend")
	default:
		fmt.Println("Harinya invalid")
	}

	// combine case
	switch dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend, yeee...")
	default:
		fmt.Println("Harinya invalid")
	}

	// switch with no expression
	var bmi = 19.0
	switch {
	case bmi < 18.5:
		fmt.Println("Anda kurus kering")
	case bmi >= 18.5 && bmi < 25.5:
		fmt.Println("Berat anda ideal")
	case bmi >= 25.5 && bmi < 30:
		fmt.Println("Anda gendut")
	default:
		fmt.Println("Anda Obesitas")
	}

	// for statement
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t", i)
	}

	fmt.Printf("\n")

	// without init statement
	j := 0
	for ; j <= 10; j++ {
		fmt.Printf("%d\t", j)
	}

	fmt.Printf("\n")

	// without increment statement
	// while loop
	k := 0
	for k <= 10 {
		fmt.Printf("%d\t", k)
		k++
	}

	fmt.Printf("\n")

	// without condition statement
	// infinite loop
	bil := 1
	for {
		if bil%2 == 0 && bil%5 == 0 {
			fmt.Printf("Bilangan yg habis dibagi 2 dan 5 adalah %d", bil)
			break
		}
		bil++
	}

	fmt.Printf("\n")
	// continue statement
	for num := 0; num < 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d\t", num)
	}
}
