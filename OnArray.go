//Linear Table:
//Sequence of elements, is a one-dimensional array.

//1. Define a one-dimensional array of student scores
//  0  1  2  3  4  5
// 90 70 50 86 60 85 => lenght = 6
package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	var lenght = len(scores)
	for i := 0; i < lenght; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
