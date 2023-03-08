package main

import "fmt"

var assign func(int, func()) int

func print() {
	fmt.Print("Called a started function")
}

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	fmt.Println("Testing the myth")
}

func test(x int, print func()) int {
	print()
	return x + 1
}

func maxNum(vals ...int) int {
	maxNum := -1

	for _, val := range vals {
		if maxNum < val {
			maxNum = val
		}
	}
	return maxNum
}

func main() {
	assign = test
	fmt.Println(assign(1, print)) // Testing Funtion Variables and passing function as an argument
	func() {                      // Testing anonymous function
		fmt.Println("Testing Anonymous functions")
	}()

	i := 1 // Testing Defer and value assignment
	defer fmt.Println(i + 1)
	i += 1
	fmt.Println(i)

	var s1 Speaker
	var d1 *Dog
	s1 = d1
	s1.Speak()
}
