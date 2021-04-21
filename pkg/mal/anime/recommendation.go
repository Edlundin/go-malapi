package anime

type Recommendation struct {
	Anime            Anime `json:"node"`
	RecommendationID int   `json:"num_recommendations"`
}
