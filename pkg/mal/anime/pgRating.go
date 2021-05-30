package anime

import (
	"encoding/json"
	"fmt"
)

//PgRating represents a PG rating..
type PgRating int

const (
	//PgRatingG represents a pg rating for all ages.
	PgRatingG PgRating = iota + 1
	//PgRatingPG represents a pg rating for children.
	PgRatingPG
	//PgRatingPG13 represents a pg rating for 13 years old and older.
	PgRatingPG13
	//PgRatingR represents a pg rating for 17 years old and older (violence & profanity).
	PgRatingR
	//PgRatingRPlus represents a pg rating for 17 years old and older (profanity & mild nudity).
	PgRatingRPlus
	//PgRatingRx represents a pg rating for adult audience (explicit nudity).
	PgRatingRx
)

var pgRatingStrDict = map[PgRating]string{
	PgRatingG:     "g",
	PgRatingPG:    "pg",
	PgRatingPG13:  "pg_13",
	PgRatingR:     "r",
	PgRatingRPlus: "r+",
	PgRatingRx:    "rx",
}

func (p *PgRating) UnmarshalJSON(b []byte) error {
	var pgRatingStr string

	if err := json.Unmarshal(b, &pgRatingStr); err != nil {
		return err
	}

	found := false

	for key, v := range pgRatingStrDict {
		if v == pgRatingStr {
			*p = key
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a PG rating", pgRatingStr)
	}

	return nil
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (p PgRating) String() string {
	pgRatingStr := "undefined"

	if str, ok := pgRatingStrDict[p]; ok {
		pgRatingStr = str
	}

	return pgRatingStr
}
