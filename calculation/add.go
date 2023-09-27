package calculation

// bisa digunakan di semua package / public
func Add(number int, numberTwo int) int {
	// return number + numberTwo
	return add(number, numberTwo)
}

// fungsi dengan capitalize kecil, hanya bisa digunakan dalam satu package
func add(number int, numberTwo int) int {
	return number + numberTwo
}
