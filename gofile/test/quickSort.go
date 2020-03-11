package main

import (
	"fmt"
)

func getIndex(array []int, low int, high int) int{
	tmp := array[low]
	for low < high {
		for low < high && array[high] >= tmp {
			high--
		}
		array[low] = array[high]
		for low < high && array[low] <= tmp {
			low++
		}
		array[high] = array[low]
	}
	array[low] = tmp
	return low
}

func quickSort(array []int, low int, high int) {
	if low < high {
		index := getIndex(array, low, high)
		quickSort(array, 0, index-1)
		quickSort(array, index+1, high)
	}
}

func main() {
	testArray := []int{ 49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22 }
	quickSort(testArray, 0, len(testArray)-1)
	fmt.Println(testArray)
}