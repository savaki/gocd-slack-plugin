package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type ApiFunc func(string, url.Values, interface{}) error

func newApiFunc(token string) ApiFunc {
	return func(method string, params url.Values, v interface{}) error {
		params.Set("token", token)

		resp, err := http.Get("https://slack.com/api/" + method + "?" + params.Encode())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(v)
	}
}
