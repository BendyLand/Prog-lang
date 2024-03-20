package main

import (
	"fmt"
	"os"
	"parser/keywords"
	"strings"
)

func main() {
	file := readFile("../../test.pr")
	file = removeComments(file)
	oneLiner := removeNewlines(file)
	withSpaces := formatString(oneLiner)
	tokens := constructTokens(withSpaces)
	test := formatTokens(tokens)
	for _, line := range test {
		fmt.Println(line)
	}
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
	for _, c := range str {
		if c == ' ' {
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
	var chars string
	for i, c := range file {
		if c == ' ' && file[i-1] == ' ' {
			continue
		}
		chars += string(c)
	}
	return string(chars)
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
