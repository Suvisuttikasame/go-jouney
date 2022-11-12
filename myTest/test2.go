package test

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

//in order to export need to capitalize first char
func Test2() {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("Hello %v \n",name)
	} else {
		log.Fatal(err)
	}
}