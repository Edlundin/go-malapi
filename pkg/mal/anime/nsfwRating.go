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

func (n *NsfwRating) UnmarshalJSON(b []byte) error {
	var nsfwRatingStr string

	if err := json.Unmarshal(b, &nsfwRatingStr); err != nil {
		return err
	}

	found := false

	for key, v := range nsfwRatingStrDict {
		if v == nsfwRatingStr {
			*n = key
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a nsfw rating", nsfwRatingStr)
	}

	return nil
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (n NsfwRating) String() string {
	nsfwRatingStr := "undefined"

	if str, ok := nsfwRatingStrDict[n]; ok {
		nsfwRatingStr = str
	}

	return nsfwRatingStr
}
