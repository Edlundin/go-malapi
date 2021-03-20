package anime

import (
	"fmt"
	"testing"
)

func Test_GenreEnum_UnmarshalJSON(t *testing.T) {
	var genre GenreEnum

	if err := genre.UnmarshalJSON([]byte("\"\"")); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := genre.UnmarshalJSON([]byte("\" \"")); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := genre.UnmarshalJSON([]byte("\"test\"")); err == nil {
		t.Errorf("failed to pass malformed list status (%q): no error returned", "test")
	}

	for k, v := range genreStrDict {
		if err := genre.UnmarshalJSON([]byte(fmt.Sprintf("%q", v))); err != nil {
			t.Errorf("failed to pass well formed list status: %s", err.Error())
		} else if genre != k {
			t.Errorf("failed to pass well formed list status: %q(%d) != %q(%d)", genre.String(), genre, v, k)
		}
	}
}
