package anime

import (
	"encoding/json"
	"fmt"
	"strings"
)

type GenreEnum int

const (
	GenreAction GenreEnum = iota + 1
	GenreAdventure
	GenreCars
	GenreComedy
	GenreDementia
	GenreDemons
	GenreDrama
	GenreEcchi
	GenreFantasy
	GenreGame
	GenreHarem
	GenreHentai
	GenreHistorical
	GenreHorror
	GenreJosei
	GenreKids
	GenreMagic
	GenreMartialArts
	GenreMecha
	GenreMilitary
	GenreMusic
	GenreMystery
	GenreParody
	GenrePolice
	GenrePsychological
	GenreRomance
	GenreSamurai
	GenreSchool
	GenreSciFi
	GenreSeinen
	GenreShoujo
	GenreShoujoAi
	GenreShounen
	GenreShounenAi
	GenreSliceOfLife
	GenreSpace
	GenreSports
	GenreSuperPower
	GenreSupernatural
	GenreThriller
	GenreVampire
	GenreYaoi
	GenreYuri
)

var genreStrDict = map[GenreEnum]string{
	GenreAction:        "action",
	GenreAdventure:     "adventure",
	GenreCars:          "cars",
	GenreComedy:        "comedy",
	GenreDementia:      "dementia",
	GenreDemons:        "demons",
	GenreDrama:         "drama",
	GenreEcchi:         "ecchi",
	GenreFantasy:       "fantasy",
	GenreGame:          "game",
	GenreHarem:         "harem",
	GenreHentai:        "hentai",
	GenreHistorical:    "historical",
	GenreHorror:        "horror",
	GenreJosei:         "josei",
	GenreKids:          "kids",
	GenreMagic:         "magic",
	GenreMartialArts:   "martial arts",
	GenreMecha:         "mecha",
	GenreMilitary:      "military",
	GenreMusic:         "music",
	GenreMystery:       "mystery",
	GenreParody:        "parody",
	GenrePolice:        "police",
	GenrePsychological: "psychological",
	GenreRomance:       "romance",
	GenreSamurai:       "samurai",
	GenreSchool:        "school",
	GenreSciFi:         "sci-fi",
	GenreSeinen:        "seinen",
	GenreShoujo:        "shoujo",
	GenreShoujoAi:      "shoujo ai",
	GenreShounen:       "shounen",
	GenreShounenAi:     "shounen ai",
	GenreSliceOfLife:   "slice of life",
	GenreSpace:         "space",
	GenreSports:        "sports",
	GenreSuperPower:    "super power",
	GenreSupernatural:  "supernatural",
	GenreThriller:      "thriller",
	GenreVampire:       "vampire",
	GenreYaoi:          "yaoi",
	GenreYuri:          "yuri",
}

func (g *GenreEnum) UnmarshalJSON(b []byte) error {
	var genreStr string

	if err := json.Unmarshal(b, &genreStr); err != nil {
		return err
	}

	found := false
	genreStr = strings.ToLower(genreStr)

	for k, v := range genreStrDict {
		if v == genreStr {
			*g = k
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a genre", genreStr)
	}

	return nil
}

func (g GenreEnum) String() string {
	genreStr := "unknown"

	if str, ok := genreStrDict[g]; ok {
		genreStr = str
	}

	return genreStr
}

type Genre struct {
	ID   int       `json:"id"`
	Name GenreEnum `json:"name"`
}
