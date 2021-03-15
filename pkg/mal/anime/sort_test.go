package anime

import "testing"

func Test_Sort_UnmarshalJSON(t *testing.T) {
	var sortType Sort

	if err := sortType.UnmarshalJSON([]byte("\"\"")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := sortType.UnmarshalJSON([]byte("\" \"")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := sortType.UnmarshalJSON([]byte("\"test\"")); err == nil {
		t.Errorf("failed to pass malformed sorting type (%q): no error returned", "test")
	}

	if err := sortType.UnmarshalJSON([]byte("\"anime_score\"")); err != nil {
		t.Errorf("failed to pass well formed sorting type: %s", err.Error())
	} else if sortType != SortByScore {
		t.Errorf("failed to pass well formed sorting type: %q != %q", sortType.String(), SortByScore.String())
	}

	if err := sortType.UnmarshalJSON([]byte("\"anime_num_list_users\"")); err != nil {
		t.Errorf("failed to pass well formed sorting type: %s", err.Error())
	} else if sortType != SortByUserListCount {
		t.Errorf("failed to pass well formed sorting type: %q != %q", sortType.String(), SortByUserListCount.String())
	}
}
