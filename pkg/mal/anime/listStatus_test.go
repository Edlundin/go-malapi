package anime

import (
	"encoding/json"
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

	if err := json.Unmarshal([]byte(`"watching"`), &listStatus); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusWatching {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusWatching.String())
	}

	if err := json.Unmarshal([]byte(`"completed"`), &listStatus); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusCompleted {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusCompleted.String())
	}

	if err := json.Unmarshal([]byte(`"dropped"`), &listStatus); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusDropped {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusDropped.String())
	}

	if err := json.Unmarshal([]byte(`"on_hold"`), &listStatus); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusOnHold {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusOnHold.String())
	}

	if err := json.Unmarshal([]byte(`"plan_to_watch"`), &listStatus); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusPlanToWatch {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusPlanToWatch.String())
	}
}
