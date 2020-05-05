package main

import (
	"fmt"
)

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}

	arr2 := arr[2:5]
	fmt.Println("arr2:", arr2)
	fmt.Println("len,cap=", len(arr2), cap(arr2))
	arr2 = append(arr2, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110)
	fmt.Println("arr2:", arr2)
	fmt.Println("len,cap=", len(arr2), cap(arr2))
	arr3 := make([]int, 10, 15)
	fmt.Println("arr3:", arr3)
	fmt.Println("len,cap=", len(arr3), cap(arr3))

}
