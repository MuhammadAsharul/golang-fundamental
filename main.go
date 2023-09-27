package main

import (
	"fmt"
	"golang-agung/calculation"
)

func main() {
	fmt.Println("Halo Saya belajar Golang")
	result := calculation.Add(11, 2)
	fmt.Println(result)

	result1 := calculation.Multiply(11, 2)
	fmt.Println(result1)

	// var name string = "Golang"
	// atau
	name := "Gojek"
	fmt.Println(name)

	// age := 9
	// if age >= 10 {
	// 	fmt.Println("Boleh bermain game")
	// } else {
	// 	fmt.Println("Kamu belum boleh bermain game")
	// }

	// score := 71
	// var grade string

	// if score > 90 {
	// 	grade = "A"
	// } else if score > 80 {
	// 	grade = "B"
	// } else if score > 70 {
	// 	grade = "C"
	// } else {
	// 	grade = "D"
	// }

	// fmt.Println(grade)

	// number := 5
	// switch number {
	// case 1:
	// 	fmt.Println("Satu")
	// case 2:
	// 	fmt.Println("Dua")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("default")
	// }

	// for i := 0; i <= 100; i++ {
	// 	fmt.Println("Halo", i)
	// }
	// j := 1
	// for j <= 100 {
	// 	fmt.Println("Halo say", j)
	// 	j++
	// }

	title := "Golang is the best language"
	// for index, letter := range title {
	// 	if index %2 == 0 {
	// 		fmt.Println("index :",index,"Letter :", string(letter))
	// 	}
	// }

	for index, letter := range title {
		stringLetter := string(letter)
		// if stringLetter == "o" || stringLetter == "a" || stringLetter == "i" || stringLetter == "e" || stringLetter == "u" {
		// 	fmt.Println("index :", index, "letter", stringLetter)
		// }
		
		switch stringLetter {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, "letter", stringLetter)
		}
	}
}
