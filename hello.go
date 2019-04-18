package main

import (
	"fmt"

	"github.com/charlesguoooo/stringutil"
)

func main() {
	fmt.Print(stringutil.Reverse("\n!oG,olleH"))
	stringutil.Package()
	stringutil.Import()
	fmt.Println(stringutil.Add(43, 12))
	a, b := stringutil.Swap("hello", "myboy")
	fmt.Println(a, b)
	ns1, ns2 := stringutil.Split(17)
	println(ns1, ns2)
}
