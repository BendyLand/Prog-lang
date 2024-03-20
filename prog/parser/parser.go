package main

import (
	"fmt"
	"os"
	"parser/keywords"
	"strings"
	// "parser/keywords"
	// "parser/expressions"
)

func main() {
	file := readFile("../../test.pr")
	file = removeComments(file)
	oneLiner := removeNewlines(file)
	withSpaces := formatString(oneLiner)
	tokens := constructTokens(withSpaces)
	test := formatTokens(tokens)
	fmt.Println(test)
}

func isEmpty(input string) bool {
	for _, c := range input {
		if c != ' ' && c != '\t' {
			return false
		}
	}
	return true
}

//! this function wrongly removes actual input as well !!!
func formatTokens(tokens []string) []string {
	var newTokens, foundKeywords []string
	placeholder := ""
	for _, token := range tokens {
		if keywords.CheckAgainstKeywords(token) {
			if len(foundKeywords) > 0 {
				temp := make([]string, 0)
				foundKeywords = temp
				newTokens = append(newTokens, placeholder)
				placeholder = ""
			}
			foundKeywords = append(foundKeywords, token)
		}
		if len(foundKeywords) > 0 && !isEmpty(token) {
			placeholder += token + " "
		}
	}
	return newTokens
}

func constructTokens(str string) []string {
	var tokens []string
	placeholder := ""
	for i, c := range str {
		if c == ' ' && str[i-1] == ' ' {
			tokens = append(tokens, placeholder)
			placeholder = ""
		}
		if c != ' ' {
			placeholder += string(c)
		}
	}
	return tokens
}

func removeNewlines(file string) string {
	return strings.Replace(file, "\n", " ", -1)
}

func formatString(file string) string {
	chars := strings.Split(file, "")
	return strings.Join(chars, " ")
}

func removeComments(file string) string {
	lines := strings.Split(file, "\n")
	for i, line := range lines {
		idx := strings.Index(line, "//")
		if idx >= 0 {
			lines[i] = line[:idx]
		}
	}
	return strings.Join(lines, "\n")
}

func readFile(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}
	return string(result)
}
