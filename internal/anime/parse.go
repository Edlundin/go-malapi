package anime

import (
	"github.com/Edlundin/go-malapi/pkg/mal"
	"github.com/Edlundin/go-malapi/pkg/mal/anime"
)

func parseSearchAnimeResponse(responseBody []byte) ([]anime.AnimeObject, mal.Page, error) {
	return []anime.AnimeObject{}, mal.Page{}, nil
}
