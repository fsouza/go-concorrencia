package main

import (
	"fmt"
)

func sum(numbers []int) {
	var result int
	for _, n := range numbers {
		result += n
	}
	fmt.Println(result)
}

func main() {
	v := []int{1, 2, 3, 4, 5, 6}
	go sum(v)
}
