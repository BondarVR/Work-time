package main

import (
	"fmt"
	"main/app"
)

func main() {
	timeString := "8:21:15PM"
	fmt.Printf(app.WorkTime(timeString))
}
