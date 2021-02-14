package maloauth2

import "net/url"

type AccessTokenRequestData struct {
	endpoint  string
	arguments url.Values
}

func (a AccessTokenRequestData) EndPoint() string {
	return a.endpoint
}

func (a AccessTokenRequestData) Arguments() url.Values {
	return a.arguments
}
