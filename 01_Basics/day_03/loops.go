package day_03

import "fmt"

func CheckLoops() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
	}

	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

}
