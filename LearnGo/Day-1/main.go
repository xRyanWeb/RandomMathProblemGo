package main

import (
	"fmt"
	"math/rand"
)

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return b / a
}

func main() {

	num1 := rand.Intn(1 + 10)
	num2 := rand.Intn(5 + 10)

	var answer int

	op := [4]string{"+", "-", "*", "/"}

	opLength := rand.Intn(len(op))
	ops := op[opLength]

	println(num1, ops, num2)

	fmt.Scan(&answer)

	if ops == "+" && answer == add(num1, num2) {

		println("Correct That: ", num1, ops, num2, "=", add(num1, num2))
	}

	if ops == "-" && answer == sub(num1, num2) {
		println("Correct That: ", num1, ops, num2, "=", sub(num1, num2))
	}

	if ops == "*" && answer == mul(num1, num2) {
		println("Correct That: ", num1, ops, num2, "=", mul(num1, num2))
	}

	if ops == "/" && answer == div(num1, num2) {
		println("Correct That: ", num1, ops, num2, "=", div(num1, num2))
	}

}
