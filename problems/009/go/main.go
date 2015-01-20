package main

import (
	"fmt"
	"math"
)

type pythagorean struct {
	a  int
	b  int
	c  int
	ok bool
}

func (pyth *pythagorean) IsPythagoreanTriplet() bool {
	if pyth.a+pyth.b+pyth.c != 1000 {
		return false
	}
	if math.Pow(float64(pyth.a), 2)+math.Pow(float64(pyth.b), 2) == math.Pow(float64(pyth.c), 2) {
		pyth.ok = true
		return true
	}
	return false
}

func main() {
	source := make(chan *pythagorean, 1000*1000*1000)
	go func() {
		for i := 1; i < 1000; i++ {
			for j := 1; j < 1000; j++ {
				for k := 1; k < 1000; k++ {
					source <- &pythagorean{a: i, b: j, c: k}
				}
			}
			fmt.Println(i)
		}
		close(source)
	}()

	found := make(chan *pythagorean, 100)
	go func() {
		for pyth := range source {
			if pyth.IsPythagoreanTriplet() {
				found <- pyth
			}
		}
		close(found)
	}()

	select {
	case answer := <-found:
		fmt.Printf("\n\nFOUND\t%+v multi %v \n", answer, answer.a*answer.b*answer.c)
	}
}
