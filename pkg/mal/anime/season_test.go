package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Season_UnmarshalJSON(t *testing.T) {
	var season Season

	if err := json.Unmarshal([]byte(`""`), &season); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &season); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &season); err == nil {
		t.Errorf("failed to pass malformed season (%q): no error returned", "test")
	}

	for seasonEnum, seasonStr := range seasonStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", seasonStr)), &season); err != nil {
			t.Errorf("failed to pass well formed season: %s", err.Error())
		} else if season != seasonEnum {
			t.Errorf("failed to pass well formed season: %q(%d) != %q(%d)", season.String(), season, seasonStr, seasonEnum)
		}
	}
}

func Test_Season_String(t *testing.T) {
	for seasonEnum, seasonStr := range seasonStrDict {
		if seasonEnum.String() != seasonStr {
			t.Errorf("failed to pass existing season: %q != %q for enum value %d", seasonEnum.String(), seasonStr, seasonEnum)
		}
	}

	var season Season

	if season.String() != "undefined" {
		t.Error(`failed to pass un-initialized season: the returned string should be "undefined"`)
	}

	season = Season(-1)

	if season.String() != "undefined" {
		t.Error(`failed to pass undefined season: the returned string should be "undefined"`)
	}
}
