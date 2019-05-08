package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo":1, "bar":2}
	fmt.Printf("%T\n",m)
	fmt.Println(m["foo"])
}