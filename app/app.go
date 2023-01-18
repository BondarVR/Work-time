package app

import (
	"fmt"
	"time"
)

const (
	layout1 = "03:04:05PM"
	layout2 = "15:04:05"
)

func WorkTime(s string) (string, error) {
	theTime, err := time.Parse(layout1, s)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	total := theTime.Format(layout2)
	return total, err
}
