package day_03

import "fmt"

// function declaration
func HelloWorld() {
	println("Hello World")
}

// function with parameters
func Add(a, b int) {
	fmt.Println(a + b)
}

// function with return values
func AddWithReturn(a, b int) (int, int) {
	return a + b, a - b
}

func CheckFunctions() {

	// function call
	HelloWorld()

	Add(10, 20)

	println(AddWithReturn(10, 20))

}
