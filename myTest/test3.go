package test

import (
	"fmt"
	"reflect"
	"strconv"
)

func Test3() {
	//variable declaration
	var v1 string = "asfadf"
	v2 := 2
	v3 := 0.5
	fmt.Printf("%v, %v, %v\n", v1, v2, v3)

	//print type
	fmt.Println(reflect.TypeOf(v3))
	
	//casting variable
	v4 := int(v3)
	v5 := "40000"
	v6, err := strconv.Atoi(v5)
	
	v7 := 40000
	v8 := strconv.Itoa(v7)

	v9 := "1.45"
	v10, err := strconv.ParseFloat(v9, 64)

	v11 := fmt.Sprintf("%f", 1.23)

	fmt.Println(v4)
	fmt.Println(v6, err, reflect.TypeOf(v6))
	fmt.Println(v8, err, reflect.TypeOf(v8))
	fmt.Println(v10, err, reflect.TypeOf(v10))
	fmt.Println(v11)
}