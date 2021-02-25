package anime

type Season int

const (
	SeasonWinter Season = iota + 1
	SeasonSpring
	SeasonSummer
	SeasonFall
)

var seasonStrDict = map[Season]string{
	SeasonWinter: "winter",
	SeasonSpring: "spring",
	SeasonSummer: "summer",
	SeasonFall:   "fall",
}

func (season Season) String() string {
	seasonStr := "unknown"

	if str, ok := seasonStrDict[season]; ok {
		seasonStr = str
	}

	return seasonStr
}
