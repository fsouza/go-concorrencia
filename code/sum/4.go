package main

import (
	"fmt"
)

func sum(numbers []int, c chan<- int) {
	var result int
	for _, n := range numbers {
		result += n
	}
	c <- result
}

func main() {
	v := []int{1, 2, 3, 4, 5, 6}
	h := len(v) / 2
	ch := make(chan int)
	go sum(v[:h], ch)
	go sum(v[h:], ch)
	fmt.Println(<-ch + <-ch)
}
