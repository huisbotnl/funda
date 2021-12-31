package main

import (
	"log"
	"regexp"
	"strings"
)

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}


func after(value string, a string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	return value[posFirstAdjusted:]
}


func before(value string, a string) string {
	// Get substring between two strings.
	posLast := strings.Index(value, a)
	if posLast == -1 {
		return ""
	}
	return value[:posLast]
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	str = ReplaceSpecialCharacter(str)
	str = strings.Title(str)
	str = strings.ReplaceAll(str, " ", "")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	//snake = strings.ReplaceAll(snake, " ", "_")
	return strings.ToLower(snake)
}

func ReplaceSpecialCharacter(text string) string{
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(text, " ")
}
