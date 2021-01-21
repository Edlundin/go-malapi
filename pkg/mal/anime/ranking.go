package anime

type AnimeRankingType int

const (
	RankTopAnime AnimeRankingType = iota + 1
	RankTopAiring
	RankTopUpcoming
	RankTopTV
	RankTopOVA
	RankTopMovie
	RankTopSpecial
	RankByPopularity
	RankFavorited
)

var rankStrDict = map[AnimeRankingType]string{
	RankTopAnime:     "all",
	RankTopAiring:    "airing",
	RankTopUpcoming:  "upcoming",
	RankTopTV:        "tv",
	RankTopOVA:       "ova",
	RankTopMovie:     "movie",
	RankTopSpecial:   "special",
	RankByPopularity: "bypopularity",
	RankFavorited:    "favorite",
}

func (rankingType AnimeRankingType) String() string {
	rankingTypeStr := "unknown"

	if str, ok := rankStrDict[rankingType]; ok {
		rankingTypeStr = str
	}

	return rankingTypeStr
}
