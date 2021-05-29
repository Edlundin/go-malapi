package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_NsfwRating_UnmarshalJSON(t *testing.T) {
	var mediaType MediaType

	if err := json.Unmarshal([]byte(`""`), &mediaType); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &mediaType); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &mediaType); err == nil {
		t.Errorf("failed to pass malformed nsfw rating (%q): no error returned", "test")
	}

	for nsfwRatingEnum, nsfwRatingStr := range mediaTypeStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", nsfwRatingStr)), &mediaType); err != nil {
			t.Errorf("failed to pass well formed nsfw rating: %s", err.Error())
		} else if mediaType != nsfwRatingEnum {
			t.Errorf("failed to pass well formed nsfw rating: %q(%d) != %q(%d)", mediaType.String(), mediaType, nsfwRatingStr, nsfwRatingEnum)
		}
	}
}

func Test_NsfwRating_String(t *testing.T) {
	for nsfwRatingEnum, nsfwRatingStr := range mediaTypeStrDict {
		if nsfwRatingEnum.String() != nsfwRatingStr {
			t.Errorf("failed to pass existing nsfw rating: %q != %q for enum value %d", nsfwRatingEnum.String(), nsfwRatingStr, nsfwRatingEnum)
		}
	}

	var mediaType MediaType

	if mediaType.String() != "undefined" {
		t.Error(`failed to pass un-initialized nsfw rating: the returned string should be "undefined"`)
	}

	mediaType = MediaType(-1)

	if mediaType.String() != "undefined" {
		t.Error(`failed to pass undefined nsfw rating: the returned string should be "undefined"`)
	}
}
