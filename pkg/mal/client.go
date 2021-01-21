package mal

type Client struct {
	apiUrl   string
	bearer   string
	clientId string
}

func (c *Client) New() (Client, error) {
	var err error

	apiClient := Client{apiUrl: "https://api.myanimelist.net/v2", clientId: ""}

	return apiClient, err
}

func (c *Client) ApiUrl() string {
	return c.apiUrl
}

func (c *Client) Bearer() string {
	return c.bearer
}

func (c *Client) SetBearer(bearer string) {
	c.bearer = bearer
}
