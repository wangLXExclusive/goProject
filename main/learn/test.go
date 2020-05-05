// main project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func variableInitial() {
	var a int = 3
	var s string = "def"
	fmt.Println(a, s)
}
func variableShort() {
	a, s := 3, "qwe"
	fmt.Println(a, s)
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation" + op)
	}
}

func main() {
	fmt.Println("Hello World!")
	// variableInitial()
	// variableShort()
	printFile("ABC.txt")
	fmt.Println(eval(3, 4, "-"))
}
