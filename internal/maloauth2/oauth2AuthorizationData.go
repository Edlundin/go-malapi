package maloauth2

//OAuth2AuthorizationData holds the data needed to obtain an access token and a refresh token from MAL's API.
type OAuth2AuthorizationData struct {
	pkceCodeVerifier string
	authorizationURL string
	state            string
}

func (o OAuth2AuthorizationData) AuthorizationURL() string {
	return o.authorizationURL
}

func (o OAuth2AuthorizationData) CodeVerifier() string {
	return o.pkceCodeVerifier
}

func (o OAuth2AuthorizationData) State() string {
	return o.state
}
