package test

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Date struct {
	day   int
	month int
	year  int
}

// setter
// don't forget specific an address
func (date *Date) SetDate(d int, m int, y int) error {
	if d < 1 || d > 31 {
		return errors.New("day is incorrect")
	} else {
		date.day = d
	}
	if m < 1 || m > 12 {
		return errors.New("month is incorrect")
	} else {
		date.month = m
	}
	if y < 1875 || d > time.Now().Year() {
		return errors.New("year is incorrect")

	} else {
		date.year = y
	}

	return nil
}

// getter
func (date Date) GetDate() (int, int, int) {
	return date.day, date.month, date.year
}

func Test9() {
	//getter and setter
	var myDate = Date{}

	err := myDate.SetDate(10, 18, 1998)
	if err == nil {
		fmt.Println(myDate.GetDate())
	} else {
		log.Fatal(err)
	}
}
