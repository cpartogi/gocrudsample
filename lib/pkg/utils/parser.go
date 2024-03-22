package utils

import "strconv"

func Generatebool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func GenerateInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
