package anime

import (
	"encoding/json"
	"fmt"
)

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

func (s *Season) UnmarshalJSON(b []byte) error {
	var seasonStr string

	if err := json.Unmarshal(b, &seasonStr); err != nil {
		return err
	}

	found := false

	for k, v := range seasonStrDict {
		if v == seasonStr {
			*s = k
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a season", seasonStr)
	}

	return nil
}

//SeasonObject represents a season JSON object
type SeasonObject struct {
	Season Season `json:"season"`
	Year   int    `json:"year"`
}
