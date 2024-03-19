package main

import (
	"fmt"
	"os"
	"strings"
	// "parser/keywords"
	// "parser/expressions"
)

func main() {
	file := readFile("../../test.pr")
	oneLiner := removeNewlines(file)
	chars := formatString(oneLiner)
	fmt.Println(chars)
}

func removeNewlines(file string) string {
	return strings.Replace(file, "\n", " ", -1)
}

func formatString(file string) string {
	chars := strings.Split(file, "")
	return strings.Join(chars, " ")
}

func readFile(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}
	return string(result)
}
