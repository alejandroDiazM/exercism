package booking

import (
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	today := time.Now()
	t, _ := time.Parse("January 2, 2006 15:04:05", date)

	return t.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	dateTime, _ := time.Parse("Thursday, July 25, 2019 13:45:00", date)
	hour := dateTime.Hour()
	minute := dateTime.Minute()

	if hour >= 12 && hour < 18 {
		if minute != 00 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	dateTime := Schedule(date)
	weekDay := dateTime.Weekday().String()
	month := dateTime.Month().String()
	day := strconv.Itoa(dateTime.Day())
	year := strconv.Itoa(dateTime.Year())
	hour := strconv.Itoa(dateTime.Hour())
	minute := strconv.Itoa(dateTime.Minute())

	descr := ("You have an appointment on " + weekDay + ", " + month + ", " + day + ", " + year + ", at " + hour + ":" + minute)
	return descr
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
