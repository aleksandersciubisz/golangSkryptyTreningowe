package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(1990, time.January, 13,
		2, 12, 0, 0, time.UTC)
	fmt.Printf("NWO best project started at %s\n:", t)
	now := time.Now()
	fmt.Printf("time now: %s\n", now)

	fmt.Println("This month is", t.Month())
	fmt.Println("This day is", t.Day())
	fmt.Println("This year is", t.Year())

	tommorow := t.AddDate(0, 0, 1)
	fmt.Println(tommorow)
	fmt.Printf("Tommorow is %v, %v, %v, %v\n",
		tommorow.Weekday(),
		tommorow.Month(),
		tommorow.Day(),
		tommorow.Year())

	longFormat := "Monday, Januar 2, 2006"
	fmt.Println("Tommorow is", tommorow.Format(longFormat))
	fmt.Println(time.Now())

}
