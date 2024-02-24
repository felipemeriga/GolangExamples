package example

import (
	"fmt"
	"math/rand"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func Time() {
	rand.Seed(time.Now().UnixNano())
	loc, _ := time.LoadLocation("America/Los_Angeles")
	year, month, day := time.Now().In(loc).Date()

	result := time.Date(year, month, day, -1, 0, 0, 0, loc).AddDate(0, 0, 1).Add(time.Minute * time.Duration(rand.Intn(240)))

	fmt.Println(result)
}
