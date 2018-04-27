package timeZone

import (
	"fmt"
	"time"
)

func GetUnixTimeZone(unixTime int64, timeZone string) time.Time {
	tm := time.Unix(unixTime, 0)
	location, _ := time.LoadLocation(timeZone)
	t := tm.In(location)
	return t
}

func GetTimeZone(dateTime time.Time, timeZone string) time.Time {
	location, _ := time.LoadLocation(timeZone)
	t := dateTime.In(location)
	return t
}

func GetCurrenTimeZone(timeZone string) time.Time {
	location, _ := time.LoadLocation(timeZone)
	t := time.Now().In(location)
	return t
}

func GetAllCurrentTime() {
	arrTimeZone := GetArrayZone()
	for _, timeZone := range arrTimeZone {
		location, _ := time.LoadLocation(timeZone)
		t := time.Now().In(location)
		fmt.Printf("Time Zone %s: %s \n", location, t)
		//fmt.Println(t)
	}
}
