package day_04

import "fmt"

func CheckSlice() {
	slc := []int{1, 2, 3, 4, 5}
	fmt.Println(slc)

	copySlice := slc
	copySlice[0] = 10
	fmt.Println("CopySlice = ", copySlice)
	fmt.Println("OriginalSlice = ", slc)
	slc1 := make([]int, len(slc))
	fmt.Println(slc1)
	copy(slc1, slc)
	fmt.Println(slc1)
	// finding type of slc1
	fmt.Printf("type of slc1 = %T\n", slc1)

}
