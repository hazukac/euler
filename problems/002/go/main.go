package main

import (
	"fmt"
)

func main() {
	f, l, limit := 0, 1, 4000000
	sum := 0
	for {
		s := f + l
		f = l
		l = s
		if l > limit {
			break
		}
		if l%2 == 0 {
			sum += l
		}
	}
	fmt.Println(sum)
}
