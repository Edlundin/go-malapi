package anime

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_GenreEnum_UnmarshalJSON(t *testing.T) {
	var genre GenreEnum

	if err := json.Unmarshal([]byte(`""`), &genre); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &genre); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &genre); err == nil {
		t.Errorf("failed to pass malformed list status (%q): no error returned", "test")
	}

	for k, v := range genreStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", v)), &genre); err != nil {
			t.Errorf("failed to pass well formed list status: %s", err.Error())
		} else if genre != k {
			t.Errorf("failed to pass well formed list status: %q(%d) != %q(%d)", genre.String(), genre, v, k)
		}
	}
}
