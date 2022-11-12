package test

import "fmt"

//generic & struct

type myConstrain interface {
	int | float64
}

type contact struct {
	name    string
	address string
	phone   string
}

// in go, doesn't have an inheritance concept, but it has composition
type balance struct {
	bal float64
	contact
}

// declare a method relate to struct
func (b balance) getBalanceInfo() float64 {
	fmt.Printf("My name is %s, work at %s and total: %.2f \n", b.contact.name, b.contact.address, b.bal)
	return b.bal
}

// define generic in func
func myFunc[T myConstrain](x T, y T) T {
	return x + y
}

func Test8() {
	//generic won't allow to have diff type arguments even you declare in an interface
	//it will always follow the first argument type
	//ex myFunc(2.5, 4) does not allow
	v := myFunc(2.5, 3.6)
	fmt.Println(v)

	contact1 := contact{
		"John",
		"222 Wall street Rd.",
		"023435354",
	}
	user1 := balance{
		32.56,
		contact1,
	}

	u1Balance := user1.getBalanceInfo()
	fmt.Println(u1Balance)

	//change user1 bal
	user1.bal = 40.00
	u1Balance = user1.getBalanceInfo()
	fmt.Println(u1Balance)

	var user2 contact
	user2.name = "Phil"
	user2.address = "443 street Bo."
	user2.phone = "12345"

	fmt.Println(user2)

	user2.address = "new add"

	fmt.Println(user2)

}
