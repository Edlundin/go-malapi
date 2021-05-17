package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Sort_UnmarshalJSON(t *testing.T) {
	var sortType Sort

	if err := json.Unmarshal([]byte(`""`), &sortType); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &sortType); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &sortType); err == nil {
		t.Errorf("failed to pass malformed sorting type (%q): no error returned", "test")
	}

	for sortEnum, sortStr := range sortStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", sortStr)), &sortType); err != nil {
			t.Errorf("failed to pass well formed sorting type: %s", err.Error())
		} else if sortType != sortEnum {
			t.Errorf("failed to pass well formed sorting type: %q(%d) != %q(%d)", sortType.String(), sortType, sortStr, sortEnum)
		}
	}
}

func Test_Sort_String(t *testing.T) {
	for sortEnum, sortStr := range seasonStrDict {
		if sortEnum.String() != sortStr {
			t.Errorf("failed to pass existing sort: %q != %q for enum value %d", sortEnum.String(), sortStr, sortEnum)
		}
	}

	var sortType Sort

	if sortType.String() != "undefined" {
		t.Error(`failed to pass un-initialized sorting type: the returned string should be "undefined"`)
	}

	sortType = Sort(-1)

	if sortType.String() != "undefined" {
		t.Error(`failed to pass undefined sorting type: the returned string should be "undefined"`)
	}
}
