package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read() {
	type name struct {
		fname string
		lname string
	}

	var filename string
	fmt.Println("Enter the file name to read the names from: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		panic("Error reading a file")
	}
	scanner := bufio.NewScanner(file)
	names := make([]name, 0)

	for scanner.Scan() {
		filename := strings.Split(scanner.Text(), " ")
		temp := new(name)
		temp.fname = filename[0]
		temp.lname = filename[1]
		names = append(names, *temp)
	}

	for _, value := range names {
		fmt.Println("First name is:", value.fname, "Second name is:", value.lname)
	}

}
