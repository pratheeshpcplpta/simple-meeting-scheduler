package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	i, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)

	// currentTime := time.Now()
	currentTime := tm

	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	fmt.Println("HH:MM:SS : ", currentTime.Format("3:4:5"))

	fmt.Println("---------------------------------------------------")

	layout := "01/02/2006"
	t, err := time.Parse(layout, "02/28/2016 03:10:20")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())

	layout = "03:04:05"
	t, err = time.Parse(layout, "04:15:58")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())

	// fmt.Println("Current Time in String: ", currentTime.String())

	// fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))

	// fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))

	// fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))

	// fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))

	// fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))

	// fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))

	// fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))

	// fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))

	// fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))

	// fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))

	// fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))

	// fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))

	// fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))

	// fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))

	// fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))

	// fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))

	// fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))
}
