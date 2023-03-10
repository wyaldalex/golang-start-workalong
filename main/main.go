package main

import "fmt"

var x int

type person struct {
	fname string
	lname string
}

//method attached to person type
func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning "`)
}

type secrentAgent struct {
	identity          person
	licenseToSabotage bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x := 7
	fmt.Println("Some variable adas", x)
	fmt.Printf("%T", x)

	xi := []int{1, 3, 91, 5, 511, 21, 31}
	fmt.Println(xi)
	fmt.Println(xi[1])
	fmt.Println("%T", xi)

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

	ops1 := secrentAgent{
		p1,
		true,
	}

	fmt.Println(p1.lname)
	fmt.Println(p1.fname)
	fmt.Printf("%T", p1)
	fmt.Println("----------------")
	p1.speak()
	fmt.Println("----------------")
	fmt.Println("-------Secret-Agent---------")
	fmt.Println(ops1.identity.fname,
		ops1.identity.lname,
		"has license to sabotage?", ops1.licenseToSabotage)
	ops1.identity.speak()

	fmt.Println("-------Human-Interface---------")
	saySomething(p1)
	saySomething(ops1.identity)

}
