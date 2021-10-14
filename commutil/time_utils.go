package commutil

import (
	"errors"
	"github.com/xkniu/go-commons-lang/core"
	"math"
	"time"
)

var TimeUtils *timeUtils

func init() {
	instance := new(timeUtils)
	TimeUtils = core.InitializeIfNeeded(instance).(*timeUtils)
}

type timeUtils struct {
}

// ---------------------------------------------------------------------
// -- New time instance
// New instance by year.
func (t *timeUtils) OfYear(year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.Now().Location())
}

// New instance by year and month.
func (t *timeUtils) OfYearMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
}

// New instance by year, month and day.
func (t *timeUtils) OfYmd(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())
}

// New instance by year, month, day and hour.
func (t *timeUtils) OfYmdh(year, month, day, hour int) time.Time {
	return time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.Now().Location())
}

// New instance by year, month, day, hour and minute.
func (t *timeUtils) OfYmdhm(year, month, day, hour, minute int) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Now().Location())
}

// New instance by year, month, day, hour, minute and second.
func (t *timeUtils) OfYmdhms(year, month, day, hour, minute, second int) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Now().Location())
}

// New instance by epoch second.
func (t *timeUtils) OfEpochSecond(epochSecond int64) time.Time {
	return time.Unix(epochSecond, 0)
}

// ---------------------------------------------------------------------
// -- Get current time in specific unit from epoch
// Get current time in seconds (from epoch).
func (t *timeUtils) CurrentTimeSeconds() int64 {
	return time.Now().Unix()
}

// Get current time in nanoseconds (from epoch).
func (t *timeUtils) CurrentTimeNanos() int64 {
	return time.Now().UnixNano()
}

// Get current time in milliseconds (from epoch).
func (t *timeUtils) CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// ---------------------------------------------------------------------
// -- Add interval to a time
// Adds a number of years to a date returning a new object.
func (t *timeUtils) AddYears(tm time.Time, amount int) time.Time {
	return tm.AddDate(amount, 0, 0)
}

// Adds a number of months to a date returning a new object.
func (t *timeUtils) AddMonths(tm time.Time, amount int) time.Time {
	return tm.AddDate(0, amount, 0)
}

// Adds a number of days to a date returning a new object.
func (t *timeUtils) AddDays(tm time.Time, amount int) time.Time {
	return tm.AddDate(0, 0, amount)
}

// Adds a number of hours to a date returning a new object.
func (t *timeUtils) AddHours(tm time.Time, amount int) time.Time {
	return tm.Add(time.Duration(amount) * time.Hour)
}

// Adds a number of minutes to a date returning a new object.
func (t *timeUtils) AddMinutes(tm time.Time, amount int) time.Time {
	return tm.Add(time.Duration(amount) * time.Minute)
}

// Adds a number of seconds to a date returning a new object.
func (t *timeUtils) AddSeconds(tm time.Time, amount int) time.Time {
	return tm.Add(time.Duration(amount) * time.Second)
}

// ---------------------------------------------------------------------
// -- Truncate time to specific field
// Truncates a date to year level returning a new object.
func (t *timeUtils) TruncateToYear(tm time.Time) time.Time {
	return time.Date(tm.Year(), time.January, 1, 0, 0, 0, 0, tm.Location())
}

// Truncates a date to month level returning a new object.
func (t *timeUtils) TruncateToMonth(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, tm.Location())
}

// Truncates a date to day level returning a new object.
func (t *timeUtils) TruncateToDay(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
}

// Truncates a date to hour level returning a new object.
func (t *timeUtils) TruncateToHour(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 0, 0, 0, tm.Location())
}

// Truncates a date to minute level returning a new object.
func (t *timeUtils) TruncateToMinute(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, tm.Location())
}

// Truncates a date to second level returning a new object.
func (t *timeUtils) TruncateToSecond(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, tm.Location())
}

// ---------------------------------------------------------------------
// -- Parse time from string
// Alias for <code>ParseDateInLocal</code>. Parse time in local time zone by layouts in order.
// @throws error if there is none layouts matched the value
func (t *timeUtils) ParseDate(value string, layouts []string) time.Time {
	return t.ParseDateInLocal(value, layouts)
}

// Parse time in local time zone by layouts in order.
// @throws error if there is none layouts matched the value
func (t *timeUtils) ParseDateInLocal(value string, layouts []string) time.Time {
	return t.ParseDateInLocation(value, time.Local, layouts)
}

// Parse time in UTC time zone by layouts in order.
// @throws error if there is none layouts matched the value
func (t *timeUtils) ParseDateInUtc(value string, layouts []string) time.Time {
	return t.ParseDateInLocation(value, time.UTC, layouts)
}

// Parse time in location time zone by layouts in order.
// @param value the value string to parse
// @param loc the time zone location to parse
// @param layouts the layouts to parse in order
// @throws error if there is none layouts matched the value
func (t *timeUtils) ParseDateInLocation(value string, loc *time.Location, layouts []string) time.Time {
	for _, layout := range layouts {
		tm, err := time.ParseInLocation(layout, value, loc)
		if err == nil {
			return tm
		}
	}
	panic(errors.New("parse date error"))
}

// ---------------------------------------------------------------------
// -- Time Comparison
// Checks if two date objects are on the same year.
func (t *timeUtils) IsSameYear(tm1 time.Time, tm2 time.Time) bool {
	return tm1.Year() == tm2.Year()
}

// Checks if two date objects are on the same year and month.
func (t *timeUtils) IsSameMonth(tm1 time.Time, tm2 time.Time) bool {
	return t.IsSameYear(tm1, tm2) && tm1.Month() == tm2.Month()
}

// Checks if two date objects are on the same year, month and day ignoring time.
func (t *timeUtils) IsSameDay(tm1 time.Time, tm2 time.Time) bool {
	return t.IsSameMonth(tm1, tm2) && tm1.Day() == tm2.Day()
}

// ---------------------------------------------------------------------
// -- Duration between time
// Return months between start and end, negative if end is before start.
// Each bound is included, eg. (2020-01-01 00:00, 2020-02-01 00:00) will return 1.
// Duration not completed to one would be ignored.
// Eg. (2020-01-01 09:00, 2020-02-01 00:00) will return 0; (2020-01-01 09:00, 2020-02-01 09:00) will return 1.
func (t *timeUtils) MonthsBetweenInclusive(start time.Time, end time.Time) int {
	if start.After(end) {
		return -1 * t.monthsBetweenInclusiveAbs(start, end)
	}
	return t.monthsBetweenInclusiveAbs(start, end)
}

// Return months between start and end, the result is always positive.
// Each bound is included, eg. (2020-01-01 00:00, 2020-02-01 00:00) will return 1.
// Duration not completed to one would be ignored.
// Eg. (2020-01-01 09:00, 2020-02-01 00:00) will return 0; (2020-01-01 09:00, 2020-02-01 09:00) will return 1.
func (t *timeUtils) monthsBetweenInclusiveAbs(start time.Time, end time.Time) int {
	if start.Equal(end) {
		return 0
	}

	var earlier time.Time // earlier time in start and end
	var later time.Time   // later time in start and end
	if start.Before(end) {
		earlier = start
		later = end
	} else {
		earlier = end
		later = start
	}

	yearDiff := later.Year() - earlier.Year()
	monthDiff := int(later.Month() - earlier.Month())

	earlierDayWithTime := time.Date(0, 0, earlier.Day(), earlier.Hour(), earlier.Minute(), earlier.Second(), earlier.Nanosecond(), earlier.Location())
	laterDayWithTime := time.Date(0, 0, later.Day(), later.Hour(), later.Minute(), later.Second(), later.Nanosecond(), later.Location())
	var diffAdjust int
	if earlierDayWithTime.After(laterDayWithTime) {
		diffAdjust = -1
	}

	return yearDiff*12 + monthDiff + diffAdjust
}

// Return days between start and end, negative if end is before start.
// Each bound is included, eg. (2020-01-01 00:00, 2020-01-02 00:00) will return 1.
// Duration not completed to one would be ignored, eg. (2020-01-01 09:00, 2020-01-02 00:00) will return 0.
func (t *timeUtils) DaysBetweenInclusive(start time.Time, end time.Time) int {
	dayDiff := end.Sub(start).Hours() / 24
	return int(math.Trunc(dayDiff))
}

// Return hours between start and end, negative if end is before start.
// Each bound is included, eg. (2020-01-01 00:00, 2020-01-01 01:00) will return 1.
// Duration not completed to one would be ignored, eg. (2020-01-01 00:00, 2020-01-01 00:59) will return 0.
func (t *timeUtils) HoursBetweenInclusive(start time.Time, end time.Time) int {
	hourDiff := end.Sub(start).Hours()
	return int(math.Trunc(hourDiff))
}
