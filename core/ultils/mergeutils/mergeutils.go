package mergeutils

import (
	"strings"
)

func PickLongerStr(str1, str2 string) string {
	if len(str1) > len(str2) {
		return str1
	}
	return str2
}

func FirstNonNil[T any](pointers ...*T) *T {
	for _, ptr := range pointers {
		if ptr != nil {
			return ptr
		}
	}
	return nil
}

func PickLongerSlice[T any](slice1, slice2 []T) []T {
	if len(slice1) > len(slice2) {
		return slice1
	}
	return slice2
}

func ToUpper(input string) string {
	return strings.ToUpper(input)
}

func ToLower(input string) string {
	return strings.ToLower(input)
}

func RemoveRedundantSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

func CapitalizeFirstLetters(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}
