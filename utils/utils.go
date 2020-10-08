package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ToSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func SearchInParsedData(inputMap []map[string]interface{}, inputWord string) {
	for i := 0; i < len(inputMap); i++ {
		if v, found := inputMap[i][inputWord]; found {
			fmt.Printf("%10v", v)
			if i%10 == 0 {
				fmt.Printf("\n")
			}
		}
	}
}
