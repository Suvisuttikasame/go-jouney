package test

import (
	"os"
	"strconv"
	"log"
	"fmt"
	"bufio"
)

func Test6() {
	//file IO
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	//close file when this func is done
	defer f.Close()
	iArr := []int {1,2,3,4,5}
	var sArr []string

	for _, v := range iArr {
		sArr = append(sArr, strconv.Itoa(v))
	}
	for _, v := range sArr {
		_, err := f.WriteString(v + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	//then open file and log in a console
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("detail :", scanner.Text())
	}
}