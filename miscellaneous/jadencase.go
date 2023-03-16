package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {

	string_slice := strings.Split(str, " ")
	cased_str := ""
	for i, v := range string_slice {
		cased_str = cased_str + strings.Title(v)
		if i != len(string_slice)-1 {
			cased_str = cased_str + " "
		}
	}
	return cased_str
}

func main() {
	fmt.Println(ToJadenCase("most trees are blue --"))
}
