package main

import "fmt"

func main() { 
	//array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	//slices
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	//slice ranges
	rangeSlice := names[1:2] // This will include elements at index 1 and 2
	fmt.Println(rangeSlice)
	
	

}