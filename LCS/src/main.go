package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	//Let strings of characters
	s := " tofoodie" // column
	t := " toody"    // row

	//Create match matrix

	matrix := make([][]int, len(t))

	for i := 0; i < len(t); i++ {
		matrix[i] = make([]int, len(s))
	}

	/*
		D[i,j] = 0    		if s[i]<>t[j]
		D[i-1, j-1] + 1   	 if s[i] == t[j]
	*/

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			fmt.Printf("|%s, %s|", string(t[i]), string(s[j]))

			if j == 0 && i == 0 {
				continue
			}

			if t[i] == s[j] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			}

		}
	}
	fmt.Println()
	fmt.Println(matrix)

}
