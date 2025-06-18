package main 

import "fmt"

func main () {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1[1])
	
	//decalre n intial arr
	arr := [3]int{1, 2, 3} 

	arr[0] = 10    // Set the first element to 10
fmt.Println(arr[0]) 


//mutidimensional arr
var matrix [2][3]int // A 2x3 matrix (2 rows, 3 columns)
matrix[0][0] = 1
fmt.Println(matrix[0][0])

 }