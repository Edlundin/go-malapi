package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Source_UnmarshalJSON(t *testing.T) {
	var source Source

	if err := json.Unmarshal([]byte(`""`), &source); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &source); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &source); err == nil {
		t.Errorf("failed to pass malformed source (%q): no error returned", "test")
	}

	for sourceEnum, sourceStr := range sourceStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", sourceStr)), &source); err != nil {
			t.Errorf("failed to pass well formed source: %s", err.Error())
		} else if source != sourceEnum {
			t.Errorf("failed to pass well formed source: %q(%d) != %q(%d)", source.String(), source, sourceStr, sourceEnum)
		}
	}
}

func Test_Source_String(t *testing.T) {
	for sourceEnum, sourceStr := range sourceStrDict {
		if sourceEnum.String() != sourceStr {
			t.Errorf("failed to pass existing source: %q != %q for enum value %d", sourceEnum.String(), sourceStr, sourceEnum)
		}
	}

	var source Source

	if source.String() != "undefined" {
		t.Error(`failed to pass un-initialized source: the returned string should be "undefined"`)
	}

	source = Source(-1)

	if source.String() != "undefined" {
		t.Error(`failed to pass undefined source: the returned string should be "undefined"`)
	}
}
