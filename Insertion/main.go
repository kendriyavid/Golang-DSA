package main

import (
	"fmt"
)

func Insertion(arr *[]int) {
	array := *arr
	length := len(array)
	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				temp := array[j]
				array[j] = array[j-1]
				array[j-1] = temp
			}
		}
	}
	fmt.Println(array)
}

func main() {

	arr := []int{7, 6, 5, 4, 3, 2, 1}
	Insertion(&arr)
}
