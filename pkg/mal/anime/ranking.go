package anime

//RankingType represents the type of anime ranking.
type RankingType int

const (
	//RankingTopAnime represents the top animes.
	RankingTopAnime RankingType = iota + 1
	//RankingTopAiring represents the top airines animes.
	RankingTopAiring
	//RankingTopUpcoming represents the top upcoming animes.
	RankingTopUpcoming
	//RankingTopTV represents the top TV animes.
	RankingTopTV
	//RankingTopOVA represents the top OVA animes.
	RankingTopOVA
	//RankingTopMovie represents the top movie animes.
	RankingTopMovie
	//RankingTopSpecial represents the top special animes.
	RankingTopSpecial
	//RankingByPopularity represents the top animes by popularity.
	RankingByPopularity
	//RankingFavorited represents the top favorited by user animes.
	RankingFavorited
)

var rankStrDict = map[RankingType]string{
	RankingTopAnime:     "all",
	RankingTopAiring:    "airing",
	RankingTopUpcoming:  "upcoming",
	RankingTopTV:        "tv",
	RankingTopOVA:       "ova",
	RankingTopMovie:     "movie",
	RankingTopSpecial:   "special",
	RankingByPopularity: "bypopularity",
	RankingFavorited:    "favorite",
}

func (rankingType RankingType) String() string {
	rankingTypeStr := "unknown"

	if str, ok := rankStrDict[rankingType]; ok {
		rankingTypeStr = str
	}

	return rankingTypeStr
}
