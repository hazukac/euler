package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type King struct {
	R int
	S string
}

func (k *King) Evaluate(str string) {
	r := 1
	for _, s := range str {
		i, _ := strconv.Atoi(string(s))
		r = r * i
	}
	if k.R < r {
		k.R = r
		k.S = str
	}
}

func main() {
	king := new(King)
	for _, str := range getSource() {
		king.Evaluate(str)
	}
	fmt.Printf("%+v\n", king)
}

func getSource() []string {
	file, err := os.Open("./source")
	if err != nil {
		log.Fatalln(err.Error())
	}
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err.Error())
	}
	all := strings.Join(strings.Split(string(buf), "\n"), "")

	res := []string{}
	for i := 0; i < len(all)-12; i++ {
		res = append(res, all[i:i+13])
	}
	return res
}
