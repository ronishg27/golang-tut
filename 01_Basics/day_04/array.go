package day_04

import "fmt"

func CheckArray() {

	arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr, "length =", len(arr))
	// fmt.Println(arr, "cap =", cap(arr))

	arrCopy := arr
	arrCopy[0] = 10
	fmt.Println("CopyArray = ", arrCopy)
	fmt.Println("OriginalArray = ", arr)

	// for i := 0; i < len(arr); i++ {
	// 	println(arr[i])
	// }
}
