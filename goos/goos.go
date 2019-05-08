package main

import (
	"fmt"
	"os"
)

func main() {
	number := 20
	// var number int =20
	// var number =20
	lol := len(os.Args[1]) + len(os.Args[0])
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(number)
	fmt.Println(lol)
}
