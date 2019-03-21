package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	age        int
	quotes     []string
}

func main() {
	person_one := person{
		first_name: "ramu",
		last_name:  "kaka",
		age:        32,
		quotes:     []string{"khana khalo", "apun naukar nahi hai"},
	}
	fmt.Println(person_one) // prints full structure for p1

	person_two := person{
		first_name: "misra",
		last_name:  "mishra",
		age:        26,
		quotes:     []string{"misra nam hai misra", "surname mishra hai vaise"},
	}
	fmt.Println(person_two) // prints full structure for p2

	array_of_persons := []person{person_one, person_two}

	for _, i := range array_of_persons {
		fmt.Println(i.first_name, i.last_name) // printing names of p1 then..... printing names of p2
		for _, j := range i.quotes {
			fmt.Println(j) // p1 quotes then...... p2 quotes then
		}
	}
}
