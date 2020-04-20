# Dates

[!["Build status"](https://api.travis-ci.com/roelofjan-elsinga/dates.svg?branch=master)](https://travis-ci.com/roelofjan-elsinga/dates)

## Installation

You can install this package in your project by running:

```bash
go get github.com/roelofjan-elsinga/dates
```

and then use it in your project like so:

```go
import "github.com/roelofjan-elsinga/dates"

func someFunctionName() {
    time, err := dates.DateStringToTime("2020-01-01")
}
```

## Available methods

- **DateStringToTime(date string) (time.Time, error)** - Convert a "2020-01-01" string to a time.Time object
- **DateTimeStringToTime(date string) (time.Time, error)** - Convert a "2020-01-01 12:00:00" string to a time.Time object
- **GetDateThisMorning() time.Time** - Get today's date at 00:00:00
- **SetTimeToNoon(date time.Time) time.Time** - Get today's date at 12:00:00
- **SetTimeToBeginDay(date time.Time) time.Time** - Get the provided date at midnight (Morning: 00:00:00:00)
- **SetTimeToEndDay(date time.Time) time.Time** - Get the provided date at midnight (Evening: 23:59:59:99)
- **TimeStringFromTime(date time.Time) string** - Convert a time.Time object to a time string (12:00:00)
- **DateStringFromTime(date time.Time) string** - Convert a time.Time object to a date string (2020-01-01)
- **DateTimeStringFromTime(date time.Time) string** - Convert a time.Time object to a date time string (2020-01-01 12:00:00)
- **DifferenceInDays(startDate time.Time, endDate time.Time) int** - Get the difference between the given dates in days
- **DifferenceInHours(startDate time.Time, endDate time.Time) int** - Get the difference between the given dates in hours
- **IsSameOrBefore(date time.Time, comparison time.Time) bool** - Determines whether the first date is equal or before the second date
- **IsSameOrAfter(date time.Time, comparison time.Time) bool** - Determines whether the first date is equal or after the second date

## Contributions

I'm open to contributions regarding more test cases and documentation.
