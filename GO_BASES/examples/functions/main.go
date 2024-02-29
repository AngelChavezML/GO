package main

func checkNumber(number int) {
	if number > 0 {
		println("Positive")
	} else if number < 0 {
		println("Negative")
	} else {
		println("Zero")
	}
}

func main() {
	checkNumber(10)
	checkNumber(-10)
	checkNumber(0)
}
