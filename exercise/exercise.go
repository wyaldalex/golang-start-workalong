package main

import (
	"fmt"
	"strings"
)

func buildHyphenedName(composed_names []string) string {
	composed_name := ""
	for _, v := range composed_names {
		start := strings.ToUpper(string(v[0]))
		composed_name = composed_name + start + string(v[1:]) + "-"
	}
	return composed_name[0 : len(composed_name)-1]
}

func bandNameGenerator(word string) string {
	fmt.Println("Calling a function with variable " + word)
	//check if the word starts and ends with the same letter
	start := strings.ToUpper(string(word[0]))
	end := string(word[len(word)-1])
	if strings.ToLower(start) == strings.ToLower(end) {
		fmt.Println("First and last character are the same")
		return start + word[1:len(word)-1] + strings.ToLower(word)
	} else {
		if strings.Contains(word, "-") {
			composed_names := strings.Split(word, "-")
			return "The " + buildHyphenedName(composed_names)
		} else {
			return "The " + start + word[1:]
		}
	}
}

func main() {
	fmt.Println("Starting a program")
	fmt.Println(bandNameGenerator("Alaska"))
	fmt.Println(bandNameGenerator("ruby-jo"))

}
