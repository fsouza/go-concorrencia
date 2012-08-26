package main

import (
	"fmt"
)

type Duck interface {
	Quack()
}

type Chicken struct {
	name string
}

func (c Chicken) Quack() {
	fmt.Printf("%s quacking...\n", c.name)
}

func main() {
	var duck Duck
	var chicken Chicken
	duck = chicken
	duck.Quack()
}
