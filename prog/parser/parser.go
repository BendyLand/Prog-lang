package main

import (
	"fmt"
	"os"
)

func main() {
	file := readFile("../../test.pr")
	fmt.Println(file)
}

func readFile(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}
	return string(result)
}
