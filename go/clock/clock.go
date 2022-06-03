package clock

import "fmt"

//Define the Clock type here.
type Clock struct {
	minutes int
}

const MinutesPerHour = 60
const HoursPerDay = 24
const MinutesPerDay = 1440

func New(hours, minutes int) Clock {
	totalMinutes := (hours * MinutesPerHour) + minutes
	return Clock{}.AdjustMinutes(totalMinutes)
}

func (c Clock) AdjustMinutes(minutes int) Clock {
	totalMinutes := (c.minutes + minutes) % MinutesPerDay
	if totalMinutes < 0 {
		totalMinutes += MinutesPerDay
	}
	return Clock{minutes: totalMinutes}
}

func (c Clock) Add(minutes int) Clock {
	return c.AdjustMinutes(minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func (c Clock) String() string {
	hours := c.minutes / MinutesPerHour
	minutes := c.minutes % MinutesPerHour
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
