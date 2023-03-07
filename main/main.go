package main

import "fmt"

var x int

type person struct {
	Fname string
	lname string
}

func main() {
	x := 7
	fmt.Println("Some variable adas", x)
	fmt.Printf("%T", x)

	xi := []int{1, 3, 91, 5, 511, 21, 31}
	fmt.Println(xi)
	fmt.Println(xi[1])
	fmt.Printf("%T", xi)

	some_map := map[string]int{
		"key1": 12121,
		"key2": 12121,
	}
	fmt.Println(some_map)
	fmt.Println(some_map["key2"])
	fmt.Printf("%T", some_map)

	p1 := person{
		"John",
		"Locke",
	}
	fmt.Println(p1.lname)
	fmt.Println(p1.Fname)
	fmt.Printf("%T", p1)

}
