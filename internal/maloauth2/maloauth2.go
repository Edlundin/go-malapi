package maloauth2

import (
	"fmt"
	"log"
	"net/url"

	"github.com/google/uuid"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

var codeVerifier string
var codeChallenge string

func GenerateMalOAuth2AuthorizationURL(clientID string, redirectURI string) {
	const malOauth2AuthorizationEndPoint string = "https://myanimelist.net/v1/oauth2/authorize"
	const responseType string = "code"
	const codeChallengeMethod string = "plain"

	pkceGenerator, err := pkce.CreateCodeVerifier()

	if err != nil {
		log.Fatalln(fmt.Sprintf("error with pkce generator: %s", err.Error()))
	}

	codeVerifier = pkceGenerator.String()
	codeChallenge = pkceGenerator.CodeChallengePlain()
	state := uuid.New().String()

	authorizationURL, err := url.Parse(malOauth2AuthorizationEndPoint)

	if err != nil {
		log.Fatalln(fmt.Sprintf("An error occurred while parsing %s as an URL: %s", malOauth2AuthorizationEndPoint, err.Error())) //TODO: replace by a proper logger
	}

	query := authorizationURL.Query()
	query.Add("response_type", responseType)
	query.Add("client_id", clientID)
	query.Add("code_challenge", codeChallenge)
	query.Add("state", state)
	query.Add("code_challenge_method", codeChallengeMethod)

	if len(redirectURI) > 0 {
		query.Add("redirect_uri", redirectURI)
	}

	authorizationURL.RawQuery = query.Encode()

	log.Println(authorizationURL)
}

