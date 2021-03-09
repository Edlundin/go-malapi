package anime

//Season represents a season (or trimester) during which animes are airing.
type Season int

const (
	//SeasonWinter represents the winter timester.
	SeasonWinter Season = iota + 1
	//SeasonSpring represents the spring timester.
	SeasonSpring
	//SeasonSummer represents the summer timester.
	SeasonSummer
	//SeasonFall represents the fall timester.
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

//SeasonObject represents a season JSON object
type SeasonObject struct {
	Season Season `json:"season"`
	Year   int    `json:"year"`
}
