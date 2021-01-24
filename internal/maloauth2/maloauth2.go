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

func ListenForMalOAuth2Callback(port uint) <-chan map[string]string {
	const malOAuth2CallbackRoute string = "/maloauth2callback"
	const codeParamName string = "code"
	const stateParamName string = "state"
	const callbackHTTPMethod string = http.MethodGet

	httpServerHandler := http.NewServeMux()
	httpServer := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%d", port), Handler: httpServerHandler}
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

	go func() { // goroutine to stop the web server when MAL have sent the code and state back
		select {
		case <-queryParamChan:
			httpServer.Shutdown(context.Background())
		}
	}()

	go func() { //goroutine to serve a web server in a non-blocking manner
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(fmt.Sprintf("An error occurred with the web server on %s: %s", httpServer.Addr, err.Error()))
		}
	}()

	return queryParamChan
}
