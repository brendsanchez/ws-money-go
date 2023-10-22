package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
	"time"
)

func TimeZone() *time.Location {
	return time.Local
}

func ConvertToFloat(value string) *float64 {
	result, err := removeDollarChar(value)
	if err != nil {
		return nil
	}
	return &result
}

func ConvertValText(value float64) string {
	if value <= 0 {
		return ""
	}
	strNumber := strconv.FormatFloat(value, 'f', -1, 64)
	return "$" + strNumber
}

func CapitalizeWords(input string) string {
	words := strings.Fields(input)

	for i, word := range words {
		if len(word) > 0 {
			words[i] = cases.Title(language.Und).String(strings.ToLower(word))
		}
	}
	output := strings.Join(words, " ")
	return output
}

func removeDollarChar(value string) (float64, error) {
	valueText := strings.Replace(value, "$", "", 1)
	return strconv.ParseFloat(valueText, 2)
}
