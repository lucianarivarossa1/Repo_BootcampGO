package main

import "fmt"

func main() {
	array := [3]int{0, 1, 2}
	fmt.Println(len(array), cap(array))
	slice := []int{0, 1, 2}
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 7)
	slice = append(slice, 9)
}
