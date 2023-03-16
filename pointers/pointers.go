package main

import (
	"fmt"
)

func main() {

	var p *int
	i, j := 100, 2888

	p = &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(*p)

	p = &j
	*p = *p * 2
	fmt.Println(j)
	fmt.Println(i)

}
