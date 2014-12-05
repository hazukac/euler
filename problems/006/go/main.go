package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	fmt.Println(
		sumOfSquares(until),
		squareOfSum(until),
		squareOfSum(until)-sumOfSquares(until),
	)
}

func squareOfSum(until int) int64 {
	res := 0
	for i := 1; i < until+1; i++ {
		res += i
	}
	return int64(math.Pow(float64(res), float64(2)))
}

func sumOfSquares(until int) int64 {
	var res int64 = 0
	for i := 1; i < until+1; i++ {
		res += int64(math.Pow(float64(i), float64(2)))
	}
	return res
}
