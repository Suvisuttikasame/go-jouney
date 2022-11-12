package test

import (
	"fmt"
)

func print10() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1 :", i)
	}
}
func print20() {
	for i := 0; i < 20; i++ {
		fmt.Println("func2 :", i)
	}
}
func chanel(ch chan bool, n int) {
	fmt.Println("start")
	for i := 0; i < 20; i++ {
		fmt.Println("chanel :", n, i)
	}
	//send value
	fmt.Println("stop")
	ch <- true
}

func Test11() {
	//concerency
	//do many tasks in parallel

	/*//code
	go print20()
	go print10()
	*/

	//need to wait before Test11 function finish
	/*//code
	time.Sleep(2 * time.Second)
	*/

	//we need to use channel to subsequencely run routine in specific order
	//chanel gonna block code until get receive value
	//create chanel
	ch := make(chan bool)

	go chanel(ch, 1)

	//receive value from chanel this gonna wait until receive value
	if <-ch {
		go chanel(ch, 2)
		fmt.Println("receive from chanel2 ", <-ch)
	}

	//after run go routine, we can not trust the order of routine that will finish first
}
