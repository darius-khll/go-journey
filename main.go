package main

import (
	"fmt"
	"go-journey/calc"
)

type Person struct {
	string
	int
}

func main() {
	p := Person{"Naveen", 50}
	fmt.Println(p)
	calc.Sum(1, 2)
}
