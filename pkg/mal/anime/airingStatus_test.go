package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_AiringStatus_UnmarshalJSON(t *testing.T) {
	var airingStatus AiringStatus

	if err := json.Unmarshal([]byte(`""`), &airingStatus); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &airingStatus); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &airingStatus); err == nil {
		t.Errorf("failed to pass malformed airing status (%q): no error returned", "test")
	}

	for airingStatusEnum, airingStatusStr := range airingStatusStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", airingStatusStr)), &airingStatus); err != nil {
			t.Errorf("failed to pass well formed airing status: %s", err.Error())
		} else if airingStatus != airingStatusEnum {
			t.Errorf("failed to pass well formed airing status: %q(%d) != %q(%d)", airingStatus.String(), airingStatus, airingStatusStr, airingStatusEnum)
		}
	}
}

func Test_AiringStatus_String(t *testing.T) {
	for airingStatusEnum, airingStatusStr := range airingStatusStrDict {
		if airingStatusEnum.String() != airingStatusStr {
			t.Errorf("failed to pass existing airing status: %q != %q for enum value %d", airingStatusEnum.String(), airingStatusStr, airingStatusEnum)
		}
	}

	var airingStatus AiringStatus

	if airingStatus.String() != "undefined" {
		t.Error(`failed to pass un-initialized airing status: the returned string should be "undefined"`)
	}

	airingStatus = AiringStatus(-1)

	if airingStatus.String() != "undefined" {
		t.Error(`failed to pass undefined airing status: the returned string should be "undefined"`)
	}
}
