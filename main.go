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

	///Specific time zone with exact date UTF
	t, _ := time.Parse(time.RFC3339, "2013-01-13T02:00:00+00:00")
	timeZone2 := "America/Bogota"
	tResult2 := timeZone.GetTimeZone(t, timeZone2)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Time Zone %s: %s \n\n", timeZone2, tResult2)

	///Specific time zone with exact unix date UTF

	timeZone3 := "America/Bogota"
	tResult3 := timeZone.GetUnixTimeZone(1294884000, timeZone3)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Time Zone %s: %s \n\n", timeZone3, tResult3)
}
