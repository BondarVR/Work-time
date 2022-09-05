package main

import (
	"fmt"
	"log"
	"main/app"
)

func main() {
	timeString := "08:21:15PM"
	res, err := app.WorkTime(timeString)
	if err != nil {
		log.Fatalf("Error: '%v'\n", err)
	}
	fmt.Println(res)
}
