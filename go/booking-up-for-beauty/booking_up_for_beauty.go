package booking

import (
	"fmt"
	"time"
)

// link : https://yourbasic.org/golang/format-parse-string-time-date-example/
// Y m d H:I:S  ==> 2006 2 1 15:04:05 --> to UTC

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"	 // day/month/year h:i:s
	t, _ := time.Parse(layout	, date)
	fmt.Printf("Parse %s --> %s \n", date, t)

	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// panic("Please implement the HasPassed function")
	layout := "January 2, 2006 15:04:05"
	parseDate, _ := time.Parse(layout, date)

	x := time.Now().After(parseDate)
	fmt.Printf("now: %s , date: %s, parsed: %s | hasPassed %v \n", time.Now(), date, parseDate, x)

	return x
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// panic("Implement IsAfternoonAppointment")
	layout := "Monday, January 2, 2006 15:04:05"
	parseDate, _ := time.Parse(layout, date)
	isAfternoon := int(parseDate.Hour()) >= 12 && int(parseDate.Hour()) < 16

	fmt.Printf("date: %s | parseDate: %s | hour %v | isAfternoon %v \n", date, parseDate, parseDate.Hour(), isAfternoon)
	
	return isAfternoon
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
	

	fmt.Printf("input date: %s | ", date)
	layout := "1/2/2006 15:04:05"
	parseDate, _ := time.Parse(layout, date)
	fmt.Printf("parseDate: %s | ", parseDate)

	formatDate := "Monday, January 2, 2006"
	descDate := parseDate.Format(formatDate)
	fmt.Printf("descDate: %s | ", descDate)

	descTime := fmt.Sprintf("%v:%v.", parseDate.Hour(), parseDate.Minute())
	fmt.Printf("descTime: %s | \n", descTime)

	desc := fmt.Sprintf("You have an appointment on %s, at %s", descDate, descTime)

	return desc
	
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	// Output: 2020-09-15 00:00:00 +0000 UTC

	// layout := "2006/02/1"
	// stringAnniv := "2012/15/09"
	// timeAnniv, _ := time.Parse(layout, stringAnniv)

	// stringThisYearAnniv := fmt.Sprintf("%v/%v/%v", time.Now().Year(), timeAnniv.Day(), int(timeAnniv.Month()))
	// thisYearAnniv, _ := time.Parse(layout, stringThisYearAnniv)
	// fmt.Printf("string anniv : %s | timeAnniv: %s | stringThisYearAnniv: %s | thisYearAnniv: %s \n", stringAnniv, timeAnniv, stringThisYearAnniv, thisYearAnniv)

	// return thisYearAnniv

	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)

}
