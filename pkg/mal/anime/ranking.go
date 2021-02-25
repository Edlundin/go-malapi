package anime

type RankingType int

const (
	RankTopAnime RankingType = iota + 1
	RankTopAiring
	RankTopUpcoming
	RankTopTV
	RankTopOVA
	RankTopMovie
	RankTopSpecial
	RankByPopularity
	RankFavorited
)

var rankStrDict = map[RankingType]string{
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

func (rankingType RankingType) String() string {
	rankingTypeStr := "unknown"

	if str, ok := rankStrDict[rankingType]; ok {
		rankingTypeStr = str
	}

	return rankingTypeStr
}
