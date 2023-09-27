package main

import (
	"fmt"
	"golang-agung/calculation"
)

func main() {
	// println
	fmt.Println("Halo Saya belajar Golang")

	// import func from deferent package
	result := calculation.Add(11, 2)
	fmt.Println(result)

	result1 := calculation.Multiply(11, 2)
	fmt.Println(result1)

	// declare variable

	// var name string = "Golang"
	// atau
	name := "Gojek"
	fmt.Println(name)

	// if else

	// age := 9
	// if age >= 10 {
	// 	fmt.Println("Boleh bermain game")
	// } else {
	// 	fmt.Println("Kamu belum boleh bermain game")
	// }

	// switch

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

	// for
	// for i := 0; i <= 100; i++ {
	// 	fmt.Println("Halo", i)
	// }
	// j := 1
	// for j <= 100 {
	// 	fmt.Println("Halo say", j)
	// 	j++
	// }

	// for text
	// title := "Golang is the best language"
	// for index, letter := range title {
	// 	if index %2 == 0 {
	// 		fmt.Println("index :",index,"Letter :", string(letter))
	// 	}
	// }

	// for index, letter := range title {
	// 	stringLetter := string(letter)
	// 	// if stringLetter == "o" || stringLetter == "a" || stringLetter == "i" || stringLetter == "e" || stringLetter == "u" {
	// 	// 	fmt.Println("index :", index, "letter", stringLetter)
	// 	// }

	// 	switch stringLetter {
	// 	case "a", "i", "u", "e", "o":
	// 		fmt.Println("index :", index, "letter", stringLetter)
	// 	}
	// }

	// array

	// var languages [5]string
	// languages[0] = "Golang"
	// languages[1] = "Ruby"
	// languages[2] = "C++"
	// languages[3] = "PHP"
	// languages[4] = "Javascript"

	languages := [...]string{"Golang", "Ruby", "C++", "PHP", "Javascript", "Python"}
	fmt.Println(languages)
	length := len(languages)
	fmt.Println(length)

	for index, lang := range languages {
		fmt.Println("index", index, "language", lang)
	}

}
