package maloauth2

type AccessTokenRequestResponseData struct {
	TokenType    string `json:"token_type"`
	ExpireIn     int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
