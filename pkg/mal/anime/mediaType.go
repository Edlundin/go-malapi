package anime

import (
	"encoding/json"
	"fmt"
)

//MediaType represents a media type.
type MediaType int

const (
	//MediaTypeTV represents the TV media.
	MediaTypeTV MediaType = iota + 1
	//MediaTypeOVA represents the OVA media.
	MediaTypeOVA
	//MediaTypeMovie represents the movie media.
	MediaTypeMovie
	//MediaTypeSpecial represents the special media.
	MediaTypeSpecial
	//MediaTypeONA represents the ONA media.
	MediaTypeONA
	//MediaTypeMusic represents the music media.
	MediaTypeMusic
	//MediaTypeUnknown represents an unknown media type.
	MediaTypeUnknown
)

var mediaTypeStrDict = map[MediaType]string{
	MediaTypeTV:      "tv",
	MediaTypeOVA:     "ova",
	MediaTypeMovie:   "movie",
	MediaTypeSpecial: "special",
	MediaTypeONA:     "ona",
	MediaTypeMusic:   "music",
	MediaTypeUnknown: "unknown",
}

func (m *MediaType) UnmarshalJSON(b []byte) error {
	var mediaTypeStr string

	if err := json.Unmarshal(b, &mediaTypeStr); err != nil {
		return err
	}

	found := false

	for key, v := range mediaTypeStrDict {
		if v == mediaTypeStr {
			*m = key
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a media type", mediaTypeStr)
	}

	return nil
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (m MediaType) String() string {
	mediaTypeStr := "undefined"

	if str, ok := mediaTypeStrDict[m]; ok {
		mediaTypeStr = str
	}

	return mediaTypeStr
}
