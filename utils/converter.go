package utils

import "strconv"

func ConverterStrToInt(data string) int {
	i, _ := strconv.Atoi(data)
	return i
}
