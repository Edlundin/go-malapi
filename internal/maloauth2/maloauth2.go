package maloauth2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

//OAuth2AuthorizationData holds the data needed to obtain an access token and a refresh token from MAL's API.
type OAuth2AuthorizationData struct {
	pkceCodeChallenge string
	pkceCodeVerifier  string
	authorizationURL  string
	state             string
}

//GenerateMalOAuth2AuthorizationURL generates an authorization URL for MAL's API.
func GenerateMalOAuth2AuthorizationURL(clientID string, redirectURI string) (OAuth2AuthorizationData, error) {
	const malOauth2AuthorizationEndPoint string = "https://myanimelist.net/v1/oauth2/authorize"
	const responseType string = "code"
	const codeChallengeMethod string = "plain"

	authorizationData := OAuth2AuthorizationData{}
	pkceGenerator, err := pkce.CreateCodeVerifier()

	if err == nil {
		return authorizationData, fmt.Errorf("pkce generation: %s", err.Error())
	}

	authorizationData.pkceCodeVerifier = pkceGenerator.String()
	authorizationData.pkceCodeChallenge = pkceGenerator.CodeChallengePlain()
	authorizationData.state = uuid.New().String()

	authorizationURL, err := url.Parse(malOauth2AuthorizationEndPoint)

	if err != nil {
		return authorizationData, fmt.Errorf("url generation: %s", err.Error())
	}

	query := authorizationURL.Query()
	query.Add("response_type", responseType)
	query.Add("client_id", clientID)
	query.Add("code_challenge", authorizationData.pkceCodeChallenge)
	query.Add("state", authorizationData.state)
	query.Add("code_challenge_method", codeChallengeMethod)

	if len(redirectURI) > 0 {
		query.Add("redirect_uri", redirectURI)
	}

	authorizationURL.RawQuery = query.Encode()
	authorizationData.authorizationURL = authorizationURL.String()

	return authorizationData, nil
}

//ListenForMalOAuth2Callback creates a web server listening on the URI "http://host:port" (passing host and port by arguments).
//Upon receiving a MAL API callback, its arguments ("code" and "state") are parsed and returned by a map-typed channel.
//After receiving these arguments, the web server will shut itself down.
func ListenForMalOAuth2Callback(host string, listeningPort uint) <-chan map[string]string {
	const malOAuth2CallbackRoute string = "/maloauth2callback"
	const codeParamName string = "code"
	const stateParamName string = "state"
	const callbackHTTPMethod string = http.MethodGet

	httpServerHandler := http.NewServeMux()
	httpServer := &http.Server{Addr: fmt.Sprintf("%s:%d", host, listeningPort), Handler: httpServerHandler}
	queryParamChan := make(chan map[string]string)

	httpServerHandler.HandleFunc(malOAuth2CallbackRoute, func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == malOAuth2CallbackRoute {
			if request.Method == callbackHTTPMethod {
				code, state := request.URL.Query().Get(codeParamName), request.URL.Query().Get(stateParamName)

				if len(code) > 0 && len(state) > 0 {
					writer.WriteHeader(200)
					queryParamChan <- map[string]string{codeParamName: code, stateParamName: state}
				} else {
					log.Println(fmt.Sprintf("request argument missing (len(%s): %d, len(%s): %d", codeParamName, len(code), stateParamName, len(state))) //TODO: hide in production environment (development log) //TODO: replace by a proper logger
				}
			} else {
				http.Error(writer, "", http.StatusMethodNotAllowed)
				log.Println(fmt.Sprintf("wanted method: %q, got: %q", callbackHTTPMethod, request.Method)) //TODO: hide in production environment (development log) //TODO: replace by a proper logger
			}
		} else {
			http.Error(writer, "", http.StatusNotFound)
			log.Println(fmt.Sprintf("received request on %s instead of %s", request.URL.Path, malOAuth2CallbackRoute)) //TODO: hide in production environment (development log) //TODO: replace by a proper logger
		}
	})

	go func() { // goroutine stopping the web server when MAL have sent back the code and state
		select {
		case <-queryParamChan:
			httpServer.Shutdown(context.Background())
		}
	}()

	go func() { //goroutine serving a web server in a non-blocking manner
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(fmt.Sprintf("An error occurred with the web server on %s: %s", httpServer.Addr, err.Error()))
		}
	}()

	return queryParamChan
}
