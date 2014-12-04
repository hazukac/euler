package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

type Set struct {
	A int
	B int
	R int64
}

type Sets []Set

func (s Sets) Len() int {
	return len(s)
}

func (s Sets) Less(i, j int) bool {
	return s[i].R < s[j].R
}

func (s Sets) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sets) Last() Set {
	return s[len(s)-1]
}

func main() {
	res := Sets{}
	var wg sync.WaitGroup
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			wg.Add(1)
			go func(k, l int) {
				if IsPalindromic(int64(k * l)) {
					res = append(res, Set{k, l, int64(k * l)})
				}
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	sort.Sort(res)
	fmt.Printf("%+v\n", res.Last())
}

func IsPalindromic(n int64) bool {
	str := strconv.FormatInt(n, 10)
	for i, s := range str {
		if string(s) != string(str[len(str)-1-i]) {
			return false
		}
	}
	return true
}
