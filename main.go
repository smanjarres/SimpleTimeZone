package main

import (
	"fmt"
	"time"

	"github.com/smanjarres/SimpleTimeZone/timeZone"
)

func main() {
	///All the zones
	timeZone.GetAllCurrentTime()

	///Current time for specific zone
	timeZone1 := "America/Bogota"
	tResult1 := timeZone.GetCurrenTimeZone(timeZone1)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Time Zone %s: %s \n\n", timeZone1, tResult1)

	///Specific time zone with exact date
	t, _ := time.Parse("yyyyMMdd hh:mm:ss", "20180426 23:22:00")
	timeZone2 := "Europe/London"
	tResult2 := timeZone.GetTimeZone(t, timeZone2)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Time Zone %s: %s \n\n", timeZone2, tResult2)
}
