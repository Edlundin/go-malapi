package anime

import (
	"encoding/json"
	"fmt"
)

//AiringStatus represents an airing status.
type AiringStatus int

const (
	//AiringStatusFinishedAiring represents a aired status.
	AiringStatusFinishedAiring AiringStatus = iota + 1
	//AiringStatusCurrentlyAiring represents a currently airing status.
	AiringStatusCurrentlyAiring
	//AiringStatusNotYetAired represents a not aired status.
	AiringStatusNotYetAired
)

var airingStatusStrDict = map[AiringStatus]string{
	AiringStatusFinishedAiring:  "finished_airing",
	AiringStatusCurrentlyAiring: "currently_airing",
	AiringStatusNotYetAired:     "not_yet_aired",
}

func (a *AiringStatus) UnmarshalJSON(b []byte) error {
	var airingStatusStr string

	if err := json.Unmarshal(b, &airingStatusStr); err != nil {
		return err
	}

	found := false

	for key, v := range airingStatusStrDict {
		if v == airingStatusStr {
			*a = key
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not an airing status", airingStatusStr)
	}

	return nil
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (a AiringStatus) String() string {
	airingStatusStr := "undefined"

	if str, ok := airingStatusStrDict[a]; ok {
		airingStatusStr = str
	}

	return airingStatusStr
}
