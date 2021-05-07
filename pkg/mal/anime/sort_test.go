package anime

import (
	"encoding/json"
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

	if err := json.Unmarshal([]byte(`"anime_score"`), &sortType); err != nil {
		t.Errorf("failed to pass well formed sorting type: %s", err.Error())
	} else if sortType != SortByScore {
		t.Errorf("failed to pass well formed sorting type: %q != %q", sortType.String(), SortByScore.String())
	}

	if err := json.Unmarshal([]byte(`"anime_num_list_users"`), &sortType); err != nil {
		t.Errorf("failed to pass well formed sorting type: %s", err.Error())
	} else if sortType != SortByUserListCount {
		t.Errorf("failed to pass well formed sorting type: %q != %q", sortType.String(), SortByUserListCount.String())
	}
}
