package anime

import (
	"encoding/json"
	"fmt"
)

//NsfwRating represents a nsfw rating.
type NsfwRating int

const (
	//NsfwRatingWhite represents a safe rating.
	NsfwRatingWhite NsfwRating = iota + 1
	//NsfwRatingGray represents a mild nudity rating.
	NsfwRatingGray
	//NsfwRatingBlack represents an explicit content rating.
	NsfwRatingBlack
)

var nsfwRatingStrDict = map[NsfwRating]string{
	NsfwRatingWhite: "white",
	NsfwRatingGray:  "gray",
	NsfwRatingBlack: "black",
}

func (m *NsfwRating) UnmarshalJSON(b []byte) error {
	var nsfwRatingStr string

	if err := json.Unmarshal(b, &nsfwRatingStr); err != nil {
		return err
	}

	found := false

	for key, v := range nsfwRatingStrDict {
		if v == nsfwRatingStr {
			*m = key
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a media type", nsfwRatingStr)
	}

	return nil
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (m NsfwRating) String() string {
	nsfwRatingStr := "undefined"

	if str, ok := nsfwRatingStrDict[m]; ok {
		nsfwRatingStr = str
	}

	return nsfwRatingStr
}
