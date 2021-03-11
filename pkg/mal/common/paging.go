package common

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type jsonURL struct {
	url.URL
}

func (j *jsonURL) UnmarshalJSON(b []byte) error {
	var str string

	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	url, err := url.Parse(str)

	if err != nil {
		return err
	}

	if len(url.Scheme) == 0 {
		return fmt.Errorf("missing scheme")
	}

	if len(url.Host) == 0 {
		return fmt.Errorf("missing host")
	}

	j.URL = *url

	return nil
}

type Paging struct {
	PreviousURL jsonURL `json:"previous"`
	NextURL     jsonURL `json:"next"`
}

func (p Paging) String() string {
	return fmt.Sprintf("{previous: %s, next: %s}", p.PreviousURL.String(), p.NextURL.String())
}
