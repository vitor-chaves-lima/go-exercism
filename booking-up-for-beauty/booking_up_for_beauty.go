package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedDate, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic("Invalid date format")
	}

	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedDate, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		panic("Invalid date format")
	}

	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if err != nil {
		panic("Invalid date format")
	}

	return parsedDate.Hour() >= 12 && parsedDate.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	scheduledTime := Schedule(date)
	formattedScheduledTime := scheduledTime.Format("Monday, January 2, 2006, at 15:04.")
	return "You have an appointment on " + formattedScheduledTime
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
