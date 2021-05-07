package common

import (
	"encoding/json"
	"testing"
)

func Test_CalendarDate_UnmarshalJSON(t *testing.T) {
	calendarDate := CalendarDate{}

	if err := json.Unmarshal([]byte(`""`), &calendarDate); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &calendarDate); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"1958-12-48"`), &calendarDate); err == nil {
		t.Error("failed to pass malformed date (48th day of the month): no error returned")
	}

	if err := json.Unmarshal([]byte(`"1958-15-03"`), &calendarDate); err == nil {
		t.Error("failed to pass malformed date (15th month of the year): no error returned")
	}

	if err := json.Unmarshal([]byte(`"1958-12-03"`), &calendarDate); err != nil {
		t.Errorf("failed to pass well formed date: %s", err.Error())
	} else {
		if calendarDate.Year() != 1958 {
			t.Errorf("failed to pass well formed date: year: %d != %d", calendarDate.Year(), 1958)
		} else if int(calendarDate.Month()) != 12 {
			t.Errorf("failed to pass well formed date: month: %d != %d", calendarDate.Month(), 12)
		} else if calendarDate.Day() != 3 {
			t.Errorf("failed to pass well formed date: day: %d != %d", calendarDate.Day(), 3)
		} else if calendarDate.String() != "1958-12-03" {
			t.Errorf("failed to pass well formed date: day: %q != %q", calendarDate.String(), "1958-12-03")
		}
	}
}
