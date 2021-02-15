package maloauth2

//OAuth2AuthorizationData holds the data needed to validate the OAuth2 MAL's API identification
type OAuth2AuthorizationData struct {
	pkceCodeVerifier string
	authorizationURL string
	state            string
}

//AuthorizationURL returns the MAL's API authorization URL
func (o OAuth2AuthorizationData) AuthorizationURL() string {
	return o.authorizationURL
}

//CodeVerifier returns the PKCE code verifier
func (o OAuth2AuthorizationData) CodeVerifier() string {
	return o.pkceCodeVerifier
}

//State returns the authorization identifier
func (o OAuth2AuthorizationData) State() string {
	return o.state
}
