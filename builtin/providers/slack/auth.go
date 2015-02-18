package slack

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type ApiFunc func(string, map[string]string, interface{}) error

func newApiFunc(token string) ApiFunc {
	return func(method string, params map[string]string, v interface{}) error {
		values := url.Values{}
		for key, value := range params {
			if value != "" {
				values.Set(key, value)
			}
		}
		values.Set("token", token)

		resp, err := http.Get("https://slack.com/api/" + method + "?" + values.Encode())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(v)
	}
}
