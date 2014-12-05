package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalln("args needed")
	}
	until, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err.Error())
	}
	numbers := []int{}
	for i := 1; i < until+1; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(
		getCommonMultiple(numbers...),
	)
}
func getCommonMultiple(numbers ...int) int {
	res := 1
	factors := []int{}
	for _, n := range numbers {
		reduced := reduce(n, factors)
		res = res * reduced
		factors = append(factors, reduced)
	}
	return res
}

func reduce(target int, factors []int) int {
	for _, factor := range factors {
		if target%factor == 0 {
			target = target / factor
		}
	}
	return target
}
