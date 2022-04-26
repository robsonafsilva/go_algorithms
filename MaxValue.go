/*
Maximum of Integer Sequences:
60 50 95 80 70
1. Algorithmic ideas
Compare arrays[i] with arrays[i + 1], if arrays[i] > arrays[i + 1] are
exchanged. So continue until the last number, arrays[length - 1] is the
maximum.
*/
package main

import "fmt"

func max(arrays []int, length int) int {
	for i := 0; i < length-1; i++ {
		fmt.Println(arrays[i])
		if arrays[i] > arrays[i+1] { // Swap
			var temp = arrays[i]
			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
			fmt.Println(arrays)
		}
	}
	var maxValue = arrays[length-1]
	return maxValue
}

func main() {
	var scores = []int{60, 50, 95, 80, 70}
	var length = len(scores)
	var maxValue = max(scores, length)
	fmt.Printf("Max Value = %d\n", maxValue)
}
