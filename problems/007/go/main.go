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
		log.Fatalln("need args")
	}
	until, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err.Error())
	}

	primes := findPrimesUntil(int64(until), []int64{2, 3})
	fmt.Println(
		len(primes),
		primes[10000],
	)
}

func findPrimesUntil(until int64, known []int64) []int64 {
	if len(known) > 10001 {
		return known
	}
	// {{ init
	eratosthenes := map[int64]bool{}
	// prepare Sieve
	for i := 1; int64(i) < until+1; i++ {
		eratosthenes[int64(i)] = false
	}
	// punch the Sieve
	for _, p := range known {
		for i := 1; ; i++ {
			if p*int64(i) > until+1 {
				break
			}
			// mark as known
			eratosthenes[p*int64(i)] = true
		}
	}
	// }}}

	// {{{ scan
	for i := 3; int64(i) < until+1; i += 2 {
		// if alreay known
		if eratosthenes[int64(i)] == true {
			continue
		}
		// if it's not prime
		if canDevice(i, known) {
			continue
		}
		// so it's prime
		known = append(known, int64(i))
		// mark multipels of found primes as known
		for j := 1; ; j++ {
			if int64(i)*int64(j) > until+1 {
				break
			}
			eratosthenes[int64(i)*int64(j)] = true
		}
	}
	// }}}

	return findPrimesUntil(until+1000, known)
}

func canDevice(n int, primes []int64) bool {
	for _, p := range primes {
		if int64(n)%p == 0 {
			return true
		}
	}
	return false
}
