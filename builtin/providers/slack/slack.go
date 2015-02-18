package slack

type Client struct {
	slack ApiFunc
}

func New(token string) *Client {
	return &Client{
		slack: newApiFunc(token),
	}
}

type ApiTestReq struct {
	Foo string
	Err string
}

type ApiTestResp struct {
	Ok    bool              `json:"ok"`
	Error string            `json:"error,omitempty"`
	Args  map[string]string `json:"args,omitempty"`
}

func (c *Client) ApiTest(input ApiTestReq) (*ApiTestResp, error) {
	params := map[string]string{
		"error": input.Err,
		"foo":   input.Foo,
	}

	v := &ApiTestResp{}
	err := c.slack("api.test", params, v)
	return v, err
}

type AuthTestResponse struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error,omitempty"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamId string `json:"team_id"`
	UserId string `json:"user_id"`
}

func (c *Client) AuthTest() (*AuthTestResponse, error) {
	v := &AuthTestResponse{}
	err := c.slack("auth.test", nil, v)
	return v, err
}
