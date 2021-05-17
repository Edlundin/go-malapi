package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_MediaType_UnmarshalJSON(t *testing.T) {
	var mediaType MediaType

	if err := json.Unmarshal([]byte(`""`), &mediaType); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &mediaType); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &mediaType); err == nil {
		t.Errorf("failed to pass malformed media type (%q): no error returned", "test")
	}

	for mediaTypeEnum, mediaTypeStr := range mediaTypeStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", mediaTypeStr)), &mediaType); err != nil {
			t.Errorf("failed to pass well formed media type: %s", err.Error())
		} else if mediaType != mediaTypeEnum {
			t.Errorf("failed to pass well formed media type: %q(%d) != %q(%d)", mediaType.String(), mediaType, mediaTypeStr, mediaTypeEnum)
		}
	}
}

func Test_MediaType_String(t *testing.T) {
	for mediaTypeEnum, mediaTypeStr := range mediaTypeStrDict {
		if mediaTypeEnum.String() != mediaTypeStr {
			t.Errorf("failed to pass existing media type: %q != %q for enum value %d", mediaTypeEnum.String(), mediaTypeStr, mediaTypeEnum)
		}
	}

	var mediaType MediaType

	if mediaType.String() != "undefined" {
		t.Error(`failed to pass un-initialized media type: the returned string should be "undefined"`)
	}

	mediaType = MediaType(-1)

	if mediaType.String() != "undefined" {
		t.Error(`failed to pass undefined media type: the returned string should be "undefined"`)
	}
}
