package anime

type AnimeSeason int

const (
	SeasonWinter AnimeSeason = iota + 1
	SeasonSpring
	SeasonSummer
	SeasonFall
)

var seasonStrDict = map[AnimeSeason]string{
	SeasonWinter: "winter",
	SeasonSpring: "spring",
	SeasonSummer: "summer",
	SeasonFall:   "fall",
}

func (season AnimeSeason) String() string {
	seasonStr := "unknown"

	if str, ok := seasonStrDict[season]; ok {
		seasonStr = str
	}

	return seasonStr
}
