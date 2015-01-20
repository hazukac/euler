package main

import (
	"fmt"

	"github.com/otiai10/primes"
)

func main() {

	var sum int64
	for _, p := range primes.Until(2 * 1000 * 1000).List() {
		sum += p
	}
	fmt.Println(sum)
}
