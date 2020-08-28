package main

import (
	"fmt"
	"time"
)

const dateFormat = "2006/01/02"

func main() {
	birthDay := "2005/01/18"
	t, err := time.Parse(dateFormat, birthDay)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("You've been born on %s", t.Weekday())
}
