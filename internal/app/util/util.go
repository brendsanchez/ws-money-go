package util

import (
	"strconv"
	"strings"
)

func ConvertValText(value string) string {
	if value == "" {
		return ""
	}
	return "$" + strings.Replace(value, ",", ".", 1)
}

func ConvertToFloat(value string) *float64 {
	result, err := removeDollarChar(value)
	if err != nil {
		return nil
	}
	return &result
}

func removeDollarChar(value string) (float64, error) {
	valueText := strings.Replace(value, "$", "", 1)
	return strconv.ParseFloat(valueText, 2)
}
