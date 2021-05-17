package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_ListStatus_UnmarshalJSON(t *testing.T) {
	var listStatus ListStatus

	if err := json.Unmarshal([]byte(`""`), &listStatus); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &listStatus); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &listStatus); err == nil {
		t.Errorf("failed to pass malformed list status (%q): no error returned", "test")
	}

	for listStatusEnum, listStatusStr := range listStatusStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", listStatusStr)), &listStatus); err != nil {
			t.Errorf("failed to pass well formed list status: %s", err.Error())
		} else if listStatus != listStatusEnum {
			t.Errorf("failed to pass well formed list status: %q(%d) != %q(%d)", listStatus.String(), listStatus, listStatusStr, listStatusEnum)
		}
	}
}

func Test_ListStatus_String(t *testing.T) {
	for listStatusEnum, listStatusStr := range mediaTypeStrDict {
		if listStatusEnum.String() != listStatusStr {
			t.Errorf("failed to pass existing list status: %q != %q for enum value %d", listStatusEnum.String(), listStatusStr, listStatusEnum)
		}
	}

	var listStatus ListStatus

	if listStatus.String() != "undefined" {
		t.Error(`failed to pass un-initialized list status: the returned string should be "undefined"`)
	}

	listStatus = ListStatus(-1)

	if listStatus.String() != "undefined" {
		t.Error(`failed to pass undefined list status: the returned string should be "undefined"`)
	}
}
