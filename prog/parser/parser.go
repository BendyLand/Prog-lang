package main

import (
	"fmt"
	"os"
	"parser/keywords"
	"slices"
	"strings"
	"regexp"
)

func main() {
	file := readFile("../../test.pr")
	file = removeComments(file)
	lines := splitIntoLines(file)
	var matches [][]string
	for _, line := range lines {
		keywordMatch := checkPatternsAgainstLine(keywords.KeywordRegexPatterns, line)
		expressionMatch := checkPatternsAgainstLine(keywords.ExpressionRegexPatterns, line)
		slices.Reverse(keywordMatch)
		slices.Reverse(expressionMatch)
		if len(keywordMatch) > 0 {
			matches = append(matches, keywordMatch)
		}
		if len(expressionMatch) > 0 {
			matches = append(matches, expressionMatch)
		}
	}
	for _, match := range matches {
		for _, token := range match {
			fmt.Println(token)
		}
	}
}

func checkPatternsAgainstLine(patterns []*regexp.Regexp, line string) []string {
	var matches []string
	for _, pattern := range patterns {
		match := pattern.Find([]byte(line))
		if match != nil {
			matches = append(matches, string(match))
		}
	}
	return matches
}

func formatString(file string) string {
	chars := strings.Split(file, "")
	var temp []string
	for _, c := range chars {
		if c != "\n" {
			temp = append(temp, c)
		}
	}
	result := strings.Join(temp, " ")
	return result
}

func splitIntoLines(file string) []string {
	lines := strings.Split(file, "\n")
	return lines
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
