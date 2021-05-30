package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_PgRating_UnmarshalJSON(t *testing.T) {
	var pgRating PgRating

	if err := json.Unmarshal([]byte(`""`), &pgRating); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &pgRating); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &pgRating); err == nil {
		t.Errorf("failed to pass malformed pg rating (%q): no error returned", "test")
	}

	for pgRatingEnum, pgRatingStr := range pgRatingStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", pgRatingStr)), &pgRating); err != nil {
			t.Errorf("failed to pass well formed pg rating: %s", err.Error())
		} else if pgRating != pgRatingEnum {
			t.Errorf("failed to pass well formed pg rating: %q(%d) != %q(%d)", pgRating.String(), pgRating, pgRatingStr, pgRatingEnum)
		}
	}
}

func Test_PgRating_String(t *testing.T) {
	for pgRatingEnum, pgRatingStr := range pgRatingStrDict {
		if pgRatingEnum.String() != pgRatingStr {
			t.Errorf("failed to pass existing pg rating: %q != %q for enum value %d", pgRatingEnum.String(), pgRatingStr, pgRatingEnum)
		}
	}

	var pgRating PgRating

	if pgRating.String() != "undefined" {
		t.Error(`failed to pass un-initialized pg rating: the returned string should be "undefined"`)
	}

	pgRating = PgRating(-1)

	if pgRating.String() != "undefined" {
		t.Error(`failed to pass undefined pg rating: the returned string should be "undefined"`)
	}
}
