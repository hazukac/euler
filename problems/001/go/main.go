package main

import (
	"fmt"
)

func sumMultiplesOf(factors []int, limit int) int {
	res := 0
	for i := 1; i < limit; i++ {
		res += summable(i, factors)
	}
	return res
}

func summable(subject int, factors []int) int {
	for _, factor := range factors {
		if subject%factor == 0 {
			return subject
		}
	}
	return 0
}

func main() {
	fmt.Println(
		sumMultiplesOf([]int{3, 5}, 10),
		sumMultiplesOf([]int{3, 5}, 1000),
	)
}
