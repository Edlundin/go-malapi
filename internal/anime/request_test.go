package anime

import (
	"testing"

	"github.com/edlundin/go-malapi/pkg/mal/anime"
)

func TestAnimeFieldsToQueryArguments(t *testing.T) {
	//Test empty array
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{}); len(queryArgs) > 0 {
		t.Error("failed to pass empty fields array: returned string is not empty")
	}

	//Test individual fields combination
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldID, anime.FieldTitle}); len(queryArgs) > 0 {
		const correctString string = "fields=id,title"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldID and FieldTitle: returned string is empty")
	}

	//Test grouped fields combination
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldMainPictureMedium, anime.FieldMainPictureLarge}); len(queryArgs) > 0 {
		const correctString string = "fields=main_picture{large,medium}"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldMainPictureMedium and FieldMainPictureLarge: returned string is empty")
	}

	//Test individual and grouped fields combination
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldID, anime.FieldMainPictureMedium, anime.FieldMainPictureLarge}); len(queryArgs) > 0 {
		const correctString string = "fields=id,main_picture{large,medium}"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldID, FieldMainPictureMedium and FieldMainPictureLarge: returned string is empty")
	}

	//Test grouped fields combination with mixed individual fields
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldID, anime.FieldMainPictureMedium, anime.FieldTitle, anime.FieldMainPictureLarge}); len(queryArgs) > 0 {
		const correctString string = "fields=id,main_picture{large,medium},title"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldID, FieldMainPictureMedium, FieldTitle, FieldMainPictureLarge: returned string is empty")
	}

	//Test multiple grouped fields
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldMainPictureMedium, anime.FieldMainPictureLarge, anime.FieldAlternativeTitlesEn, anime.FieldAlternativeTitlesJa}); len(queryArgs) > 0 {
		const correctString string = "fields=alternative_titles{en,ja},main_picture{large,medium}"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldMainPictureMedium, FieldMainPictureLarge, FieldAlternativeTitlesEn and FieldAlternativeTitlesJa: returned string is empty")
	}

	//Test mixed grouped fields
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldMainPictureMedium, anime.FieldAlternativeTitlesEn, anime.FieldMainPictureLarge, anime.FieldAlternativeTitlesJa}); len(queryArgs) > 0 {
		const correctString string = "fields=alternative_titles{en,ja},main_picture{large,medium}"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldMainPictureMedium, FieldAlternativeTitlesEn, FieldMainPictureLarge and FieldAlternativeTitlesJa: returned string is empty")
	}

	//Test mixed multiple individual and multiple grouped fields
	if queryArgs := animeFieldsToQueryArguments([]anime.Field{anime.FieldMainPictureMedium, anime.FieldAlternativeTitlesEn, anime.FieldID, anime.FieldTitle, anime.FieldMainPictureLarge, anime.FieldAlternativeTitlesJa}); len(queryArgs) > 0 {
		const correctString string = "fields=alternative_titles{en,ja},id,main_picture{large,medium},title"

		if queryArgs != correctString {
			t.Errorf("%q != %q", queryArgs, correctString)
		}
	} else {
		t.Error("failed to pass fields array with FieldMainPictureMedium, FieldAlternativeTitlesEn, FieldMainPictureLarge and FieldAlternativeTitlesJa: returned string is empty")
	}
}
