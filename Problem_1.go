package main

import "fmt"

func main() {
	// variable to save
	sum := 0
	for i := 1; i < 100000000; i++ {
		if i%5 == 0 {
			sum += i
		} else if i%3 == 0 {
			sum += i
		} else {
			continue
		}

	}
	fmt.Println(sum)
}
