package test

import (
	"fmt"
)

func Test4() {
	// array & slices 
	//array

	var arr1 [5]int
	arr1[0] = 3

	//declare and assign value
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr2)

	//iterate array
	for i:=0; i<len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for i, v := range arr1 {
		fmt.Println(v, i)
	}

	//multi dim array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i:=0; i<2; i++ {
		for j:=0; j<2; j++ {
			fmt.Println(arr3[i][j])
		}
	}


	aStr1 := "abdsafdf"
	rArr := []rune(aStr1)
	
	for _, v := range rArr {
		fmt.Println(v)
	}
	
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	fmt.Println(bStr)


	//slices
	s1 := make([]string, 6)
	s1[0] = "test"
	s1[1] = "demo"
	fmt.Println(s1, len(s1))

	arr4 := [5]int {1,2,3,4,5}
	s2 := arr4[0:2]
	fmt.Println(arr4[:3])
	fmt.Println(arr4[2:])
	fmt.Println(s2)
	//pointer to register when change mother array slice also change
	arr4[0] = 10
	fmt.Println(s2)
	
	//append value to slice
	s2 = append(s2, 22)
	fmt.Println(s2)
	fmt.Println(arr4)


}