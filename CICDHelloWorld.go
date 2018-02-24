package main

import (
	"fmt"
)

func main() {
	a, b, c := Content()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func Content() (string, string, string) {
	a := "Built using:"
	b := "   Git	     = Source Control"
	c := "	 Travis-CI   = Continuous Intergration"
	return a, b, c
	// Each time I add a new element I'll add a Println...

}
