package expressions

import (
	"log"
	"regexp"
)

type Expression int

const (
	Logic Expression = iota
	Math  Expression = iota
)

func createRegex(pattern string) regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error creating regex")
	}
	return *r
}
