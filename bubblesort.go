package main

import "fmt"

func Bubblesort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			Swap(a, j)
		}
	}
}

func Swap(a []int, index int) {
	if a[index] > a[index+1] {
		temp := a[index+1]
		a[index+1] = a[index]
		a[index] = temp
	}
}

func bubblesort() {

	fmt.Println("Enter a sequence of 10 integers")
	a := make([]int, 10)
	count := 0
	var num int
	for {
		if count == 10 {
			break
		}
		fmt.Scan(&num)
		a[count] = num
		count += 1
	}
	Bubblesort(a)
	fmt.Println("Sorted slice is", a)

}
