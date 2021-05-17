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
		t.Errorf("failed to pass malformed genre (%q): no error returned", "test")
	}

	for genreEnum, genreStr := range genreStrDict {
		if err := json.Unmarshal([]byte(fmt.Sprintf("%q", genreStr)), &genre); err != nil {
			t.Errorf("failed to pass well formed genre: %s", err.Error())
		} else if genre != genreEnum {
			t.Errorf("failed to pass well formed genre: %q(%d) != %q(%d)", genre.String(), genre, genreStr, genreEnum)
		}
	}
}

func Test_GenreEnum_String(t *testing.T) {
	for genreEnum, genreStr := range genreStrDict {
		if genreEnum.String() != genreStr {
			t.Errorf("failed to pass existing genre: %q != %q for enum value %d", genreEnum.String(), genreStr, genreEnum)
		}
	}

	var genre GenreEnum

	if genre.String() != "undefined" {
		t.Error(`failed to pass un-initialized genre: the returned string should be "undefined"`)
	}

	genre = GenreEnum(-1)

	if genre.String() != "undefined" {
		t.Error(`failed to pass undefined genre: the returned string should be "undefined"`)
	}
}
