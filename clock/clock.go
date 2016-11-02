// Package clock provides methods to add minutes and hours to calculate the time from midnight
package clock

import "fmt"

const testVersion = 4

// Clock is a struct of hour and minute, initialises to 0 by default
type Clock struct {
	hour   int
	minute int
}

func cleanHourMinute(hour, minute int) (h, m int) {
	remainingMinutes := 0
	// calculate minutes
	if minute >= 0 {
		if minute > 59 {
			m = minute % 60
			h = minute / 60
		} else {
			m = minute
		}
	} else {
		if minute < -59 {
			remainingMinutes = minute % -60
			m = 60 + remainingMinutes
			h = -(minute / -60)
			if remainingMinutes < 0 {
				h--
			}
		} else {
			m = 60 + minute
			h--
		}
	}

	// add remaining hours from minutes
	hour += h

	// calculate hours
	if hour >= 0 {
		if hour > 23 {
			h = hour % 24
		} else {
			h = hour
		}
	} else {
		if hour < -23 {
			h = 24 + (hour % -24)
		} else {
			h = 24 + hour
		}

	}

	// if exactly 60 mins or 24 hours then reset to 0
	if h == 24 {
		h = 0
	}
	if m == 60 {
		m = 0
	}
	return
}

// New returns a new time with the hour and minute cleaned up
func New(hour, minute int) Clock {
	h, m := cleanHourMinute(hour, minute)

	a := Clock{
		hour:   h,
		minute: m,
	}

	return a
}

// String zero pads if necessary the hour and time in a readable format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to existing time
func (c Clock) Add(minutes int) Clock {
	newMinutes := c.minute + minutes
	c.hour, c.minute = cleanHourMinute(c.hour, newMinutes)
	return c
}
