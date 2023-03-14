package main

import (
	"fmt"
)

func cycle_some(x []int) {
	for i, v := range x {
		fmt.Println("Index value: ", i, " Value: ", v)
	}
}

func cycle_some_map(x map[string]int) {

	for k, v := range x {
		fmt.Println("key value: ", k, " Value: ", v)
	}
}

type employee interface {
	get_id()
}

type engineer struct {
	speciality int
	code       string
}

func (emp engineer) get_id() {
	fmt.Println(emp.code, "-", emp.speciality)
}

func print_employee_id(emp employee) {
	emp.get_id()
}

func main() {
	fmt.Println("---------------Slice testing-------------------")
	x := []int{7, 8, 9, 10}
	fmt.Println("Accesing value in slice", x[2])

	//Slices are mutable
	y := make([]int, 0, 50) //created fixed size slice
	y = append(y, 101)
	y = append(y, 200)
	fmt.Println("Slices are mutable", y)

	x = append(x, 122121)
	cycle_some(x)

	fmt.Println("---------------Maps Testing-------------------")
	x_map := map[string]int{"John Locke": 56, "Sarah Conor": 43, "Ellen Ripley": 38}
	fmt.Println("Accesing an element: ", x_map["Sarah Conor"])
	cycle_some_map(x_map)

	y_map := make(map[string]string)
	y_map["John"] = "Locke"

	fmt.Println(y_map)

	fmt.Println("---------------Structs Testing-------------------")
	engineer_1 := engineer{14, "CSE401"}
	print_employee_id(engineer_1)
}
