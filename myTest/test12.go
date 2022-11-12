package test

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	balance int
	//to allow user to use this account only one times
	lock sync.Mutex
}

func (a *account) getAccount() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *account) withDraw(c int) bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	if c < a.balance {
		fmt.Println("account", a.balance)
		a.balance = a.balance - c
		fmt.Println("remain", a.balance)
		return true
	} else {
		fmt.Println("not enough money")
		return false
	}

}

func Test12() {
	//mutex & lock
	//in order to lock while someone is using method in type
	var myAccount = account{}
	myAccount.balance = 100

	fmt.Println("get balance ", myAccount.getAccount())

	for i := 1; i < 20; i++ {
		go myAccount.withDraw(10)
	}

	time.Sleep(2 * time.Second)

}
