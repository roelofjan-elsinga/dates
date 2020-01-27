package dates

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateStringToTimeReturnsNoErrors(t *testing.T) {

	_, err := DateStringToTime("2020-01-01")

	require.NoError(t, err, "Expected no error")
}

func TestDateTimeStringToTimeReturnsNoErrors(t *testing.T) {

	_, err := DateTimeStringToTime("2020-01-01 12:00:00")

	require.NoError(t, err, "Expected no error")
}

func TestGetDateThisMorningReturnsTodayAtMidnight(t *testing.T) {

	now := time.Now()
	time := GetDateThisMorning()

	assert.Equal(t, now.Year(), time.Year(), "The year should be the same")
	assert.Equal(t, now.Month(), time.Month(), "The month should be the same")
	assert.Equal(t, now.Day(), time.Day(), "The day should be the same")
	assert.Equal(t, 0, time.Hour(), "The hour should be 0")
	assert.Equal(t, 0, time.Minute(), "The minute should be 0")
	assert.Equal(t, 0, time.Second(), "The second should be 0")
	assert.Equal(t, 0, time.Nanosecond(), "The nano second should be 0")
}

func TestSetTimeToNoonReturnsTodayAtNoon(t *testing.T) {

	now := time.Now()
	time := SetTimeToNoon(now)

	assert.Equal(t, now.Year(), time.Year(), "The year should be the same")
	assert.Equal(t, now.Month(), time.Month(), "The month should be the same")
	assert.Equal(t, now.Day(), time.Day(), "The day should be the same")
	assert.Equal(t, 12, time.Hour(), "The hour should be 12")
	assert.Equal(t, 0, time.Minute(), "The minute should be 0")
	assert.Equal(t, 0, time.Second(), "The second should be 0")
	assert.Equal(t, 0, time.Nanosecond(), "The nano second should be 0")
}

func TestSetTimeToBeginDayReturnsTodayAtMidnight(t *testing.T) {

	now := time.Now()
	time := SetTimeToBeginDay(now)

	assert.Equal(t, now.Year(), time.Year(), "The year should be the same")
	assert.Equal(t, now.Month(), time.Month(), "The month should be the same")
	assert.Equal(t, now.Day(), time.Day(), "The day should be the same")
	assert.Equal(t, 0, time.Hour(), "The hour should be 0")
	assert.Equal(t, 0, time.Minute(), "The minute should be 0")
	assert.Equal(t, 0, time.Second(), "The second should be 0")
	assert.Equal(t, 0, time.Nanosecond(), "The nano second should be 0")
}

func TestSetTimeToEndDayReturnsLastNanoSecondToday(t *testing.T) {

	now := time.Now()
	time := SetTimeToEndDay(now)

	assert.Equal(t, now.Year(), time.Year(), "The year should be the same")
	assert.Equal(t, now.Month(), time.Month(), "The month should be the same")
	assert.Equal(t, now.Day(), time.Day(), "The day should be the same")
	assert.Equal(t, 23, time.Hour(), "The hour should be 23")
	assert.Equal(t, 59, time.Minute(), "The minute should be 59")
	assert.Equal(t, 59, time.Second(), "The second should be 59")
	assert.Equal(t, 99, time.Nanosecond(), "The nano second should be 99")
}

func TestDifferenceInDaysReturnsTheCorrectNumber(t *testing.T) {

	startDate, _ := DateStringToTime("2020-01-01")
	endDate, _ := DateStringToTime("2020-01-08")

	time := DifferenceInDays(startDate, endDate)

	assert.Equal(t, 7, time, "The difference should be 7")
}

func TestDifferenceInHoursReturnsTheCorrectNumber(t *testing.T) {

	startDate, _ := DateTimeStringToTime("2020-01-01 12:00:00")
	endDate, _ := DateTimeStringToTime("2020-01-01 17:00:00")

	time := DifferenceInHours(startDate, endDate)

	assert.Equal(t, 5, time)
}

func TestRoundFloatToIntReturnsRoundedOffInteger(t *testing.T) {

	roundDownValue := RoundFloatToInt(3.4)

	assert.Equal(t, 3, roundDownValue, "The rounded number should be 3")
	assert.Equal(t, 4, RoundFloatToInt(3.5), "The rounded number should be 4")
	assert.Equal(t, 4, RoundFloatToInt(3.6), "The rounded number should be 4")
}

func TestGetDateTimeStringFromTimeReturnsCorrectDateTimeString(t *testing.T) {

	dateTime := "2020-01-01 12:00:00"

	time, _ := DateTimeStringToTime(dateTime)

	date := GetDateTimeStringFromTime(time)

	assert.Equal(t, dateTime, date)
}

func TestGetDateStringFromTimeReturnsCorrectDateTimeString(t *testing.T) {

	dateTime := "2020-01-01 12:00:00"

	time, _ := DateTimeStringToTime(dateTime)

	date := GetDateStringFromTime(time)

	assert.Equal(t, "2020-01-01", date)
}

func TestGetTimeStringFromTimeReturnsCorrectDateTimeString(t *testing.T) {

	dateTime := "2020-01-01 12:00:00"

	time, _ := DateTimeStringToTime(dateTime)

	date := GetTimeStringFromTime(time)

	assert.Equal(t, "12:00:00", date)
}
