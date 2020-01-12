package utils

import (
	"strconv"
)

func StringToFloat64(str string) float64 {

	var result float64

	result, _ = strconv.ParseFloat(str, 64)

	return result
}
