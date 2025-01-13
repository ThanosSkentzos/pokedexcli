package main

import "strings"


func cleanInput(text string) []string{
	text = strings.ToLower(text)
	result := strings.FieldsFunc(text,Separators)

	return result
}

func Separators(r rune) bool{
	return strings.ContainsRune(" ,.",r)
}