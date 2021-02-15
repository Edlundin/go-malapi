package maloauth2

import "net/url"

//AccessTokenRequestData holds the data needed to obtain an access token and a refresh token from MAL's API.
type AccessTokenRequestData struct {
	endpoint  string
	arguments url.Values
}

//EndPoint returns the access token request endpoint
func (a AccessTokenRequestData) EndPoint() string {
	return a.endpoint
}

//Arguments returns the access token request arguments (grant_type, client_id, client_secret, code, code_verifier)
func (a AccessTokenRequestData) Arguments() url.Values {
	return a.arguments
}
