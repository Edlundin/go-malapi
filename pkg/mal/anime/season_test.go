package anime

import (
	"encoding/json"
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

	if err := json.Unmarshal([]byte(`"winter"`), &season); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonWinter {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonWinter.String())
	}

	if err := json.Unmarshal([]byte(`"fall"`), &season); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonFall {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonFall.String())
	}

	if err := json.Unmarshal([]byte(`"summer"`), &season); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonSummer {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonSummer.String())
	}

	if err := json.Unmarshal([]byte(`"spring"`), &season); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonSpring {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonSpring.String())
	}
}
