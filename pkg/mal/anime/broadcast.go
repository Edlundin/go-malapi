package anime

import (
	"encoding/json"
	"fmt"
	"time"
)

type jsonWeekday struct {
	time.Weekday
}

func (j jsonWeekday) String() string {
	return j.Weekday.String()
}

func (j *jsonWeekday) UnmarshalJSON(b []byte) error {
	var str string

	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	switch str {
	case "monday":
		j.Weekday = time.Monday
	case "tuesday":
		j.Weekday = time.Tuesday
	case "wednesday":
		j.Weekday = time.Wednesday
	case "thursday":
		j.Weekday = time.Thursday
	case "friday":
		j.Weekday = time.Friday
	case "saturday":
		j.Weekday = time.Saturday
	case "sunday":
		j.Weekday = time.Sunday
	default:
		return fmt.Errorf("%q is not a day", str)
	}

	return nil
}

type jsonTime struct {
	time.Time
}

func (j jsonTime) String() string {
	return fmt.Sprintf("%d:%d", j.Hour(), j.Minute())
}

func (j *jsonTime) UnmarshalJSON(b []byte) error {
	var str string

	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	parsedTime, err := time.Parse("15:04", str)

	if err != nil {
		return err
	}

	j.Time = parsedTime

	return nil
}

type Broadcast struct {
	DayOfTheWeek jsonWeekday `json:"day_of_the_week"`
	StartTime    jsonTime    `json:"start_time"` //only Hour and Minutes are set
}
