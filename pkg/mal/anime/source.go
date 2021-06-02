package anime

import (
	"encoding/json"
	"fmt"
)

//Source represents an original work that inspired its adaptation.
type Source int

const (
	SourceOriginal Source = iota + 1
	SourceManga
	Source4KomaManga
	SourceWebManga
	SourceDigitalManga
	SourceNovel
	SourceLightNovel
	SourceVisualNovel
	SourceGame
	SourceCardGame
	sourceBook
	sourcePictureBook
	SourceRadio
	SourceMusic
	SourceOther
)

var sourceStrDict = map[Source]string{
	SourceOriginal:     "original",
	SourceManga:        "manga",
	Source4KomaManga:   "4_koma_manga",
	SourceWebManga:     "web_manga",
	SourceDigitalManga: "digital_manga",
	SourceNovel:        "novel",
	SourceLightNovel:   "light_novel",
	SourceVisualNovel:  "visual_novel",
	SourceGame:         "game",
	SourceCardGame:     "card_game",
	sourceBook:         "book",
	sourcePictureBook:  "picture_book",
	SourceRadio:        "radio",
	SourceMusic:        "music",
	SourceOther:        "other",
}

func (a *Source) UnmarshalJSON(b []byte) error {
	var sourceStr string

	if err := json.Unmarshal(b, &sourceStr); err != nil {
		return err
	}

	found := false

	for sourceEnum, sourceEnumStr := range sourceStrDict {
		if sourceEnumStr == sourceStr {
			*a = sourceEnum
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a valid source", sourceStr)
	}

	return nil
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (a Source) String() string {
	sourceStr := "undefined"

	if str, ok := sourceStrDict[a]; ok {
		sourceStr = str
	}

	return sourceStr
}
