/*
Bubble Sorting Algorithm:
Compare arrays[j] with arrays[j + 1], if arrays[j] > arrays[j + 1] are
exchanged.
Remaining elements repeat this process, until sorting is completed.
*/
package main

import "fmt"

func main() {
	//index starts from 0
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Println("Before Sort: ", scores)
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(arrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arrays[j] > arrays[j+1] { //swap
				var flag = arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = flag

			}
		}
	}
}
