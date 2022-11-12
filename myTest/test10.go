package test

import (
	"fmt"
)

// define struct of objects
type game struct {
	name               string
	price              float64
	sastificationPoint int
	rate               float64
}

type book struct {
	name               string
	price              float64
	sastificationPoint int
	rate               float64
}

// define an interface
// in order to check which type sastify this interface
// type need to contain checkSaleRate method
type checkRating interface {
	checkSaleRate() int
}

// assign checkSaleRate method
func (g game) checkSaleRate() int {
	return int(g.rate) * int(g.sastificationPoint)
}
func (b book) checkSaleRate() int {
	return int(b.rate) / int(b.sastificationPoint)
}

func Test10() {
	//interface
	//it is a contract[protocol] to declare the method that every type should have
	//and it is an abstract type
	//summary the interface is just checking point of type, which methods or properties it should contain.

	var myGame checkRating
	myGame = game{
		"overwatch",
		99.99,
		4,
		10.11,
	}

	var myBook checkRating
	myBook = book{
		"comic",
		10.44,
		3,
		134.23,
	}

	fmt.Println(myGame.checkSaleRate())
	fmt.Println(myBook.checkSaleRate())

}
