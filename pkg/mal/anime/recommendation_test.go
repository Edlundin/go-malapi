package anime

import (
	"encoding/json"
	"testing"
)

func Test_Recommendation_UnmarshalJSON(t *testing.T) {
	singleRecommendation := `{"node":{"id":1535,"title":"Death Note","main_picture":{"medium":"https://api-cdn.myanimelist.net/images/anime/9/9453.jpg","large":"https://api-cdn.myanimelist.net/images/anime/9/9453l.jpg"}},"num_recommendations":561}`

	recommendation := Recommendation{}

	if err := json.Unmarshal([]byte(`""`), &recommendation); err == nil {
		t.Error("failed to pass empty string: no error returned")
	}

	if err := json.Unmarshal([]byte(`" "`), &recommendation); err == nil {
		t.Error("failed to pass space string: no error returned")
	}

	if err := json.Unmarshal([]byte(`"test"`), &recommendation); err == nil {
		t.Errorf("failed to pass malformed weekday status (%q): no error returned", "test")
	}

	if err := json.Unmarshal([]byte(singleRecommendation), &recommendation); err != nil {
		t.Errorf("failed to pass correct recommendation: %s", err.Error())
	}

	if recommendation.RecommendationID != 561 {
		t.Errorf("failed to pass correct recommendation: recommendation ID different from %d (%d)", 561, recommendation.RecommendationID)
	}

	if recommendation.Anime.ID != 1535 {
		t.Errorf("failed to pass correct recommendation: anime ID different from %d (%d)", 1535, recommendation.Anime.ID)
	}

	if recommendation.Anime.Title != "Death Note" {
		t.Errorf("failed to pass correct recommendation: anime title different from %s (%s)", "Death Note", recommendation.Anime.Title)
	}

	if recommendation.Anime.MainPicture.Medium != "https://api-cdn.myanimelist.net/images/anime/9/9453.jpg" {
		t.Errorf("failed to pass correct recommendation: anime main picture medium different from %s (%s)", "https://api-cdn.myanimelist.net/images/anime/9/9453.jpg", recommendation.Anime.MainPicture.Medium)
	}

	if recommendation.Anime.MainPicture.Large != "https://api-cdn.myanimelist.net/images/anime/9/9453l.jpg" {
		t.Errorf("failed to pass correct recommendation: anime main picture medium different from %s (%s)", "https://api-cdn.myanimelist.net/images/anime/9/9453l.jpg", recommendation.Anime.MainPicture.Large)
	}
}
