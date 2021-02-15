package maloauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

//GenerateMalOAuth2AuthorizationURL generates an authorization URL for MAL's API
func GenerateMalOAuth2AuthorizationURL(clientID string, redirectURI string) (OAuth2AuthorizationData, error) {
	const malOauth2AuthorizationEndPoint string = "https://myanimelist.net/v1/oauth2/authorize"
	const responseType string = "code"
	const codeChallengeMethod string = "plain"

	authorizationData := OAuth2AuthorizationData{}
	pkceGenerator, err := pkce.CreateCodeVerifier()

	if err != nil {
		return authorizationData, fmt.Errorf("pkce generation: %s", err.Error())
	}

	pkceCodeChallenge := pkceGenerator.CodeChallengePlain()
	authorizationData.pkceCodeVerifier = pkceGenerator.String()
	authorizationData.state = uuid.New().String()

	endpointURL, err := url.Parse(malOauth2AuthorizationEndPoint)

	if err != nil {
		return authorizationData, fmt.Errorf("endpoint URL parsing: %s", err.Error())
	}

	query := endpointURL.Query()
	query.Add("response_type", responseType)
	query.Add("client_id", clientID)
	query.Add("code_challenge", pkceCodeChallenge)
	query.Add("state", authorizationData.state)
	query.Add("code_challenge_method", codeChallengeMethod)

	if len(redirectURI) > 0 {
		query.Add("redirect_uri", redirectURI)
	}

	endpointURL.RawQuery = query.Encode()
	authorizationData.authorizationURL = endpointURL.String()

	return authorizationData, nil
}

//GenerateMalOAuth2AccessTokenRequest generates the data needed to form an HTTP POST request to obtain the access token
func GenerateMalOAuth2AccessTokenRequest(clientID string, clientSecret string, code string, codeVerifier string) AccessTokenRequestData {
	const malOAuth2AccessTokenEndPoint string = "https://myanimelist.net/v1/oauth2/token"
	const grantType string = "authorization_code"

	requestArguments := url.Values{}

	requestArguments.Add("grant_type", grantType)
	requestArguments.Add("client_id", clientID)
	requestArguments.Add("client_secret", clientSecret)
	requestArguments.Add("code", code)
	requestArguments.Add("code_verifier", codeVerifier)

	return AccessTokenRequestData{
		endpoint:  malOAuth2AccessTokenEndPoint,
		arguments: requestArguments,
	}
}

//ParseMalOAuth2AccessTokenResponse parses the body of an access token request's response and returns the parsed data or an error
func ParseMalOAuth2AccessTokenResponse(httpResponse []byte, httpResponseStatusCode int) (AccessTokenRequestResponseData, error) {
	var responseData AccessTokenRequestResponseData

	if httpResponseStatusCode == http.StatusOK {
		err := json.Unmarshal(httpResponse, &responseData)

		if err != nil {
			return responseData, fmt.Errorf("demarshalling JSON response body (code: %d): %s", httpResponseStatusCode, err.Error())
		}
	} else {
		const errorTypeKey string = "error"
		const errorMessageKey string = "message"
		var responseErrorData map[string]interface{}

		err := json.Unmarshal(httpResponse, &responseErrorData)

		if err != nil {
			return responseData, fmt.Errorf("demarshalling JSON response body (code: %d): %s", httpResponseStatusCode, err.Error())
		}

		requestErrorType, ok := responseErrorData[errorTypeKey]

		if !ok {
			return responseData, fmt.Errorf("getting MAL's API error type (code: %d): %q key not found", httpResponseStatusCode, errorTypeKey)
		}

		requestErrorMessage, ok := responseErrorData[errorMessageKey]

		if !ok {
			return responseData, fmt.Errorf("getting MAL's API error message (code: %d): %q key not found", httpResponseStatusCode, errorMessageKey)
		}

		return responseData, fmt.Errorf("request error (code: %d): error type: %q, message: %q", httpResponseStatusCode, requestErrorType, requestErrorMessage)
	}

	return responseData, nil
}

//ListenForMalOAuth2Callback creates a web server listening on the URI "http://host:port" (passing host and port by arguments).
//Upon receiving a MAL API callback, its arguments ("code" and "state") are parsed, if the state pass the check, the code is returned by a string-typed channel.
//After that, the web server will shut itself down
func ListenForMalOAuth2Callback(host string, listeningPort uint, state string) <-chan string {
	const malOAuth2CallbackRoute string = "/maloauth2callback"
	const codeParamName string = "code"
	const stateParamName string = "state"
	const callbackHTTPMethod string = http.MethodGet

	httpServerHandler := http.NewServeMux()
	httpServer := &http.Server{Addr: fmt.Sprintf("%s:%d", host, listeningPort), Handler: httpServerHandler}
	queryParamChan := make(chan string)

	httpServerHandler.HandleFunc(malOAuth2CallbackRoute, func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == malOAuth2CallbackRoute {
			if request.Method == callbackHTTPMethod {
				receivedCode, receivedState := request.URL.Query().Get(codeParamName), request.URL.Query().Get(stateParamName)

				if len(receivedCode) > 0 && state == receivedState {
					writer.WriteHeader(200)
					queryParamChan <- receivedCode
				} else {
					http.Error(writer, "", http.StatusBadRequest)
					log.Println(fmt.Sprintf("request argument missing or incorrect:\nreceivedState: %s (%d)\nwantedState: %s (%d)\nreceivedCode: %s (%d))",
						receivedState,
						len(receivedState),
						state,
						len(state),
						receivedCode,
						len(receivedCode))) //TODO: hide in production environment (development log) //TODO: replace by a proper logger
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
