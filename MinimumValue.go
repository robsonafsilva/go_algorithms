/*
1. Algorithmic ideas
Initial value minIndex=0, j=1 Compare arrays[minIndex] with arrays[j]
if arrays[minIndex] > arrays[j] then minIndex=j, j++ else j++. continue until
the last number, arrays[minIndex] is the Min Value.
*/
package main

import "fmt"

func min(arrays []int, length int) int {
	var minIndex = 0 // the index of the minimum
	for j := 1; j < length; j++ {
		if arrays[minIndex] > arrays[j] {
			minIndex = j
		}
	}
	return arrays[minIndex]
}

func main() {
	var scores = []int{60, 80, 95, 50, 70}
	var length = len(scores)
	var minValue = min(scores, length)
	fmt.Printf("Min Value = %d\n", minValue)
}
