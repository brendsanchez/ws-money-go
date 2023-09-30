package util

import (
	"fmt"
	"strconv"
	"strings"
)

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

func AddDollarChar(value string) string {
	if value == "" {
		return value
	}
	return fmt.Sprintf("%s:%s", "$", value)
}
