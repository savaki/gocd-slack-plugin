package slack

import (
	"errors"
	"net/url"

	"github.com/savaki/gocd-slack-plugin/form"
)

type Self struct {
	Team   string
	User   string
	TeamId string
	UserId string
}

type Client struct {
	Self Self
	api  ApiFunc
}

func New(token string) (*Client, error) {
	api := newApiFunc(token)

	authInfo, err := authTest(api)
	if err != nil {
		return nil, err
	}
	if !authInfo.Ok {
		return nil, errors.New(authInfo.Error)
	}

	client := &Client{
		Self: Self{
			Team:   authInfo.Team,
			TeamId: authInfo.TeamId,
			User:   authInfo.User,
			UserId: authInfo.UserId,
		},
		api: api,
	}

	return client, nil
}

type ApiTestReq struct {
	Foo string `form:"foo"`
	Err string `form:"err"`
}

type ApiTestResp struct {
	Ok    bool              `json:"ok"`
	Error string            `json:"error,omitempty"`
	Args  map[string]string `json:"args,omitempty"`
}

func (c *Client) ApiTest(input ApiTestReq) (*ApiTestResp, error) {
	values := form.AsValues(input)
	resp := &ApiTestResp{}
	err := c.api("api.test", values, resp)
	return resp, err
}

type AuthTestResponse struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error,omitempty"`
	Url    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamId string `json:"team_id"`
	UserId string `json:"user_id"`
}

func (c *Client) AuthTest() (*AuthTestResponse, error) {
	return authTest(c.api)
}

func authTest(api ApiFunc) (*AuthTestResponse, error) {
	resp := &AuthTestResponse{}
	err := api("auth.test", url.Values{}, resp)
	return resp, err
}
