package common

import (
	"encoding/json"
	"fmt"
	"time"
)

type CalendarDate struct {
	time.Time
}

func (c *CalendarDate) UnmarshalJSON(b []byte) error {
	var str string

	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02", str)
	c.Time = parsedTime

	return err
}

//String returns the date following the format YYY-MM-DD.
func (c CalendarDate) String() string {
	return fmt.Sprintf("%d-%02d-%02d", c.Year(), int(c.Month()), c.Day())
}
