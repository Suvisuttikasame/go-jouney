package test

import (
	"fmt"
)

func changePointer(mrPtr *int) {
	*mrPtr = 12
}

func changeArrPointer(arrPtr *[4]int) {
	for i:=0; i<4; i++ {
		arrPtr[i] *= 2
	}
}

func changeArr(arrPtr [4]int) {
	for i:=0; i<4; i++ {
		arrPtr[i] *= 2
	}
}

func getAverageSlice(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, v := range nums {
		sum += v
	}
	return (sum / numSize)
}

func Test5() {
	//pointer allow to access the specific location in memory
	mrInt := 10
	fmt.Println("before change:", mrInt)
	changePointer(&mrInt)
	fmt.Println("before change:", mrInt)

	//demon pointer assign address
	var f1 *int = &mrInt
	fmt.Println("f1 address", f1)
	fmt.Println("f1 value", *f1)

	//pointer array
	arrTest := [4]int {1,2,3,4}
	fmt.Println("before change", arrTest)
	changeArrPointer(&arrTest)
	fmt.Println("before change", arrTest)

	//pointer slice
	mySli := []float64{11,19,17}
	val := getAverageSlice(mySli...)
	fmt.Printf("avg float %.3f", val)

}