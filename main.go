package main

import "fmt"

func main() {
	fmt.Println("hello!")
	task1()
	fmt.Println()
	task2()
}

func task1() {
	var firstNumber int = 20
	var secondNumber int = 30

	firstNumber, secondNumber = secondNumber, firstNumber

	fmt.Print("firstNumber: ", firstNumber)
	fmt.Println()
	fmt.Print("secondNumber: ", secondNumber)

}

func task2() {
	var firstName string = "Ante"
	var lastName string = "Jolić"

	var fullname string = firstName + " " + lastName

	fmt.Println("Fullname: ", fullname)
}
