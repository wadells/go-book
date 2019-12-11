package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func printStringConcat(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(ioutil.Discard, s)
}

func printStringJoin(args []string) {
	s := strings.Join(args, " ")
	fmt.Fprintln(ioutil.Discard, s)
}
