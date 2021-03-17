package anime

import "testing"

func Test_ListStatus_UnmarshalJSON(t *testing.T) {
	var listStatus ListStatus

	if err := listStatus.UnmarshalJSON([]byte("\"\"")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := listStatus.UnmarshalJSON([]byte("\" \"")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := listStatus.UnmarshalJSON([]byte("\"test\"")); err == nil {
		t.Errorf("failed to pass malformed list status (%q): no error returned", "test")
	}

	if err := listStatus.UnmarshalJSON([]byte("\"watching\"")); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusWatching {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusWatching.String())
	}

	if err := listStatus.UnmarshalJSON([]byte("\"completed\"")); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusCompleted {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusCompleted.String())
	}

	if err := listStatus.UnmarshalJSON([]byte("\"dropped\"")); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusDropped {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusDropped.String())
	}

	if err := listStatus.UnmarshalJSON([]byte("\"on_hold\"")); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusOnHold {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusOnHold.String())
	}

	if err := listStatus.UnmarshalJSON([]byte("\"plan_to_watch\"")); err != nil {
		t.Errorf("failed to pass well formed list status: %s", err.Error())
	} else if listStatus != ListStatusPlanToWatch {
		t.Errorf("failed to pass well formed list status: %q != %q", listStatus.String(), ListStatusPlanToWatch.String())
	}
}
