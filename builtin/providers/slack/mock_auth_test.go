package slack

import (
	"encoding/json"
	"net/url"
	"strings"
)

type MockApiFunc struct {
	Method   string
	Params   url.Values
	Response string
}

func (m *MockApiFunc) Handle(method string, params url.Values, v interface{}) error {
	m.Method = method
	m.Params = params

	return json.NewDecoder(strings.NewReader(m.Response)).Decode(v)
}
