package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func init() {
	flag.Parse()
}

var subject int
var sqrt int

type KnownFactors map[int]bool

func (k KnownFactors) Init() {
	for i := 2; i < sqrt; i++ {
		k[i] = false
	}
}

func (k KnownFactors) Has(i int) bool {
	if deleted, ok := k[i]; ok && !deleted {
		return false
	}
	return true
}

func (k KnownFactors) Add(i int) {
	k[i] = true
	for j := 1; j < sqrt; j++ {
		if j*i > sqrt {
			break
		}
		fmt.Print(".")
		k[i*j] = true
	}
	fmt.Print("\n")
}

func (k KnownFactors) CanDevide(i int) bool {
	for n, deleted := range k {
		if !deleted {
			continue
		}
		if float64(n) > math.Sqrt(float64(i)) {
			break
		}
		if i%n == 0 {
			return true
		}
	}
	return false
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	var err error
	subject, err = strconv.Atoi(args[0])
	if err != nil {
		return
	}
	sqrt = int(math.Floor(math.Sqrt(float64(subject))))

	factors := []int{}
	known := KnownFactors{}
	known.Init()
	if subject%2 == 0 {
		known.Add(2)
		factors = append(factors, 2)
	}
	for i := 3; i < sqrt; i += 2 {
		if known.Has(i) {
			continue
		}
		if !known.CanDevide(i) {
			fmt.Println(i)
			known.Add(i)
			if subject%i == 0 {
				factors = append(factors, i)
			}
		}
	}
	fmt.Println(factors)
}
