package state

import (
	"fmt"
	"os"
	"regexp"
)

type LineParser struct {
	CurrentChar rune
	NextChar    rune
	Index       int
	State       string
}

func parseStringLiteral(line string) string {
	pattern, err := regexp.Compile("\".*\"")
	if err != nil {
		fmt.Println("There was a problems creating the regex")
		os.Exit(1)
	}
	return pattern.FindString(line)
}

func (lp *LineParser) Parse(line string) []string {
	var tokens []string
	for i := range len(line) - 1 {
		lp.UpdateState(i, rune(line[i]), rune(line[i+1]))
		fmt.Printf("%d %c %c %s\n", lp.Index, lp.CurrentChar, lp.NextChar, lp.State)
	}
	return tokens
}

func (lp *LineParser) UpdateState(index int, current, next rune) {
	lp.Index = index
	lp.CurrentChar = current
	lp.NextChar = next
	switch {
	case lp.CurrentChar == '"' && lp.State != "StringLiteral":
		lp.State = "StringLiteral"
	case lp.CurrentChar == '"' && lp.State == "StringLiteral":
		lp.State = "Parser"
	case lp.CurrentChar == '/' && lp.NextChar == '/':
		lp.State = "Comment"
	case lp.State == "Comment" && lp.NextChar == '\n':
		lp.State = "Parser"
	}

}
