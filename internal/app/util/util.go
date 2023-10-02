package util

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

func TimeZone() *time.Location {
	value := os.Getenv("TZ")
	if value == "" {
		value = "America/Buenos_Aires"
	}

	loc, err := time.LoadLocation(value)
	if err != nil {
		logrus.Fatal(err)
	}
	return loc
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
