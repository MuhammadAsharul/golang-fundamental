package main

import (
	"fmt"
)

// "golang-agung/calculation"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// methdod
func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, dengan email : %s", user.FirstName, user.LastName, user.Email)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, dengan email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	isAvailable bool
}

func main() {
	user1 := User{1, "Zelda", "Xavier", "xavier@gmail.com", true}
	user2 := User{2, "Zechic", "Arachio", "arachio@gmail.com", true}

	result := user1.Display()
	fmt.Println(result)
	fmt.Println(user2.Display())

	// display1 := displayUser(user1)
	// display2 := displayUser(user2)

	// fmt.Println(display1)
	// fmt.Println(display2)
	// users := []User{user1, user2}

	// group := Group{"Gamer", user1, users, true}
	// displayGroup(group)

	// fmt.Println(user)
	// fmt.Println(user.firstName)

	// println
	// fmt.Println("Halo Saya belajar Golang")

	// import func from deferent package
	// result := calculation.Add(11, 2)
	// fmt.Println(result)

	// result1 := calculation.Multiply(11, 2)
	// fmt.Println(result1)

	// declare variable

	// var name string = "Golang"
	// atau
	// name := "Gojek"
	// fmt.Println(name)

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

	// languages := [...]string{"Golang", "Ruby", "C++", "PHP", "Javascript", "Python"}
	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	// for index, lang := range languages {
	// 	fmt.Println("index", index, "language", lang)
	// }

	// slice
	// var gamingSlice []string
	// gamingSlice = append(gamingSlice, "Playstation 4")
	// gamingSlice = append(gamingSlice, "Nintendo Switch")
	// gamingSlice = append(gamingSlice, "XBox")
	// fmt.Println(gamingSlice)

	// map
	// myMap := map[string]string{
	// 	"Golang": "Golang is Best language",
	// 	"Ruby":   "Ruby is the best framework",
	// 	"Python": "Hello World",
	// }
	// fmt.Println(myMap)
	// for _, value := range myMap {
	// 	// fmt.Println("Key : ", key, "Value :", value)
	// 	fmt.Println("Value:", value)
	// }
	// fmt.Println("=====")
	// delete(myMap, "Golang")
	// fmt.Println(myMap)

	// slice of map
	// students := []map[string]string{
	// 	{"name": "Arul", "score": "A"},
	// 	{"name": "Aziz", "score": "B"},
	// 	{"name": "Abi", "score": "C"},
	// }
	// for _, value := range students {
	// 	fmt.Println(value["name"], " - ", value["score"])
	// }

	// kuis
	// 1. hitung rata-rata
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// var total int
	// for _, score := range scores {
	// 	total = total + score
	// }

	// length := len(scores)
	// average := float64(total) / float64(length)
	// fmt.Println(average)

	// 2. append to slice
	// var goodScores []int

	// for _, score := range scores {
	// 	if score >= 90 {
	// 		goodScores = append(goodScores, score)
	// 	}
	// }
	// fmt.Println(goodScores)

	// function
	// sentence := printMyResult("Python")
	// fmt.Println(sentence)

	// result := add(10, 12)
	// fmt.Println(result)

	// luas, _ := calculate(10, 12)
	// fmt.Println("Luas persegi panjang : ", luas)
	// fmt.Println(keliling)

	// quiz
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Println(total)

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	// result, err := calculate(10, 2, "/")
	// 	result, err := calculate(10, 2, "=")
	// 	if err != nil {
	// 		fmt.Println("Terjadi Kesalahan")
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(result)
	// }

}

// func printMyResult(sentence string) string {
// 	// newSentence := sentence + " Belajar"
// 	return sentence
// }
// func add(number, numberTwo int) int {
// 	return number + numberTwo
// }

// func calculate(panjang int, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)
// 	return
// }

// quiz
// func sum(numbers []int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func calculate(number, numberTwo int, operation string) (int, error) {
// 	var result int
// 	var errorResult error
// 	switch operation {
// 	case "+":
// 		result = number + numberTwo
// 	case "-":
// 		result = number - numberTwo
// 	case "*":
// 		result = number * numberTwo
// 	case "/":
// 		result = number / numberTwo
// 	default:
// 		errorResult = errors.New("Unknown Operation")
// 	}
// 	return result, errorResult
// }

// struct sbg parameter (function)
// func displayUser(user User) string {
// 	return fmt.Sprintf("Name : %s %s dengan email : %s", user.firstName, user.lastName, user.Email)
// }

// embedded struct
// func displayGroup(group Group) {
// 	fmt.Printf("Name : %s", group.Name)
// 	fmt.Printf("\nMember Count : %d \n", len(group.Users))
// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }
