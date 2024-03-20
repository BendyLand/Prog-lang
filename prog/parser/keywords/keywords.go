package keywords

import (
	"fmt"
	"regexp"
)

var keywordPatterns = []string{
	"print\\s+",
	"puts\\s+",
	"let\\s+",
	"let\\s+mut\\s+=\\s*|\\n",
	"if\\s+",
	"elif\\s+",
	"else\\s+",
	"for\\s+",
	"return\\s+",
}

var KeywordRegexPatterns = compilePatternsToRegexp(keywordPatterns)

func compilePatternsToRegexp(patterns []string) []*regexp.Regexp {
	var result []*regexp.Regexp
	for _, pattern := range patterns {
		temp, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println(pattern)
			fmt.Println("Error compiling regex:", err)
			continue
		}
		result = append(result, temp)
	}
	return result
}
