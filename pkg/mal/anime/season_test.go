package anime

import "testing"

func Test_Season_UnmarshalJSON(t *testing.T) {
	var season Season

	if err := season.UnmarshalJSON([]byte("\"\"")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := season.UnmarshalJSON([]byte("\" \"")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := season.UnmarshalJSON([]byte("\"test\"")); err == nil {
		t.Errorf("failed to pass malformed season (%q): no error returned", "test")
	}

	if err := season.UnmarshalJSON([]byte("\"winter\"")); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonWinter {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonWinter.String())
	}

	if err := season.UnmarshalJSON([]byte("\"fall\"")); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonFall {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonFall.String())
	}

	if err := season.UnmarshalJSON([]byte("\"summer\"")); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonSummer {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonSummer.String())
	}

	if err := season.UnmarshalJSON([]byte("\"spring\"")); err != nil {
		t.Errorf("failed to pass well formed season: %s", err.Error())
	} else if season != SeasonSpring {
		t.Errorf("failed to pass well formed season: %q != %q", season.String(), SeasonSpring.String())
	}
}
