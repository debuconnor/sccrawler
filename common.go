package sccrawler

import (
	"strconv"
	"time"
)

func Timestamp() string {
	currentTime := time.Now()
	year := strconv.Itoa(currentTime.Year())
	month := currentTime.Month().String()
	day := strconv.Itoa(currentTime.Day())
	hour := strconv.Itoa(currentTime.Hour())
	minute := strconv.Itoa(currentTime.Minute())
	second := strconv.Itoa(currentTime.Second())

	convertTime := "[" + year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second + "] "
	return convertTime
}

func checkError(e error, comment string) {
	if e != nil {
		panic(Timestamp() + comment + "\n" + e.Error() + "\n")
	}
}
