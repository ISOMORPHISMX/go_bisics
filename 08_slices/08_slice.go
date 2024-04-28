package main

import (
	"fmt"
	//"slices"
)

func main() {

	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "first"
	s[1] = "second"
	s[2] = "third"
	//s[3] = "fourth"

	fmt.Println("set: ", s )

	s = append(s, "append1")
	s = append(s, "append2", "append3")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := c[2:5]
	fmt.Println("lists: ", l)

}