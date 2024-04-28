package main

import "fmt"

func main() {
	i := 1 
	for i < 3 {
		fmt.Println("i = ", i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println("j = ", j)
	}

	for {
		fmt.Println("loop ")
		fmt.Println("i = ", i)
		if i > 0 {
			break
		}
	}


}