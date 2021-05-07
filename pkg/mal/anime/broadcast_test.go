package anime

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_jsonWeekday_UnmarshalJSON(t *testing.T) {
	var jWeekday jsonWeekday

	if err := json.Unmarshal([]byte(`""`), &jWeekday); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &jWeekday); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &jWeekday); err == nil {
		t.Errorf("failed to pass malformed weekday status (%q): no error returned", "test")
	}

	if err := json.Unmarshal([]byte(`"monday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Monday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Monday.String())
	}

	if err := json.Unmarshal([]byte(`"tuesday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Tuesday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Tuesday.String())
	}

	if err := json.Unmarshal([]byte(`"wednesday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Wednesday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Wednesday.String())
	}

	if err := json.Unmarshal([]byte(`"thursday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Thursday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Thursday.String())
	}

	if err := json.Unmarshal([]byte(`"friday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Friday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Friday.String())
	}

	if err := json.Unmarshal([]byte(`"saturday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Saturday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Saturday.String())
	}

	if err := json.Unmarshal([]byte(`"sunday"`), &jWeekday); err != nil {
		t.Errorf("failed to pass well formed weekday status: %s", err.Error())
	} else if jWeekday.Weekday != time.Sunday {
		t.Errorf("failed to pass well formed weekday status: %q != %q", jWeekday.String(), time.Sunday.String())
	}
}

func Test_jsonTime_UnmarshalJSON(t *testing.T) {
	var jTime jsonTime

	if err := json.Unmarshal([]byte(`""`), &jTime); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &jTime); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"25:30"`), &jTime); err == nil {
		t.Error("failed to pass malformed time (25th hour of the day): no error returned")
	}

	if err := json.Unmarshal([]byte(`"12:65"`), &jTime); err == nil {
		t.Error("failed to pass malformed time (65th minute of the hour): no error returned")
	}

	if err := json.Unmarshal([]byte(`"12:30"`), &jTime); err != nil {
		t.Errorf("failed to pass well formed time: %s", err.Error())
	} else if jTime.Hour() != 12 {
		t.Errorf("failed to pass well formed time: year: %d != %d", jTime.Hour(), 12)
	} else if jTime.Minute() != 30 {
		t.Errorf("failed to pass well formed time: month: %d != %d", jTime.Minute(), 30)
	} else if jTime.String() != "12:30" {
		t.Errorf("failed to pass well formed time: day: %q != %q", jTime.String(), "12:30")
	}
}
