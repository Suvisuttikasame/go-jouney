package test

import (
	"fmt"
)

func Test7() {
	//map
	//the structure of obj, which represent key value pair

	//basic declaration
	//1. var myMap map [keyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)

	//2. use make to declare
	villians := make(map[string]string)

	//assign values
	heroes["batman"] = "bruce wayne"
	heroes["spiderman"] = "peter parker"

	villians["Joker"] = "Joker"
	villians["Thanos"] = "Thanos"

	fmt.Println("hero", heroes)
	fmt.Println("villian", villians)

	//delete value in map
	delete(heroes, "spiderman")

	fmt.Println(heroes)

}
