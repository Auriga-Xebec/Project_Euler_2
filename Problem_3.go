//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

//firstly lets find factors of a given number

package main

import (
	"fmt"
)

func main() {
	loops := 13195
	var factors []int

	for i := 1; i < loops; i++ {
		if loops%i == 0 {
			factors = append(factors, i)
		} else {
			continue
		}

	}
	for i, v := range factors {
		fmt.Println(i)
		for x := 0; x < v; x++ {
			fmt.Println(x)
			fmt.Println(v)
		}
	}
	//fmt.Println(factors)
}
