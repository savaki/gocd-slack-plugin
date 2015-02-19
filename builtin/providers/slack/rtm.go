package slack

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type RtmStartResp struct {
	Ok            bool      `json:"ok"`
	Error         string    `json:"error,omitempty"`
	Url           string    `json:"url"`
	Self          User      `json:"self"`
	Team          Team      `json:"team"`
	LatestEventTs string    `json:"latest_event_ts"`
	Channels      []Channel `json:"channels"`
	Groups        []Group   `json:"groups"`
	Users         []User    `json:"users"`
	Bots          []Bot     `json:"bots"`
	CacheVersion  string    `json:"cache_version"`
}

func (c *Client) RtmStart() (*RtmStartResp, error) {
	resp := &RtmStartResp{}
	err := c.api("rtm.start", url.Values{}, resp)
	return resp, err
}

func (c *Client) Listen(handler interface{}) error {
	resp, err := c.RtmStart()
	if err != nil {
		return err
	}

	u, err := url.Parse(resp.Url)
	if err != nil {
		return err
	}

	host := u.Host
	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:443", host)
	}
	rawConn, err := tls.Dial("tcp", host, nil)
	if err != nil {
		return err
	}

	wsHeaders := http.Header{
		"Origin":                   {resp.Url},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}

	wsConn, _, err := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
	if err != nil {
		return err
	}

	for {
		err = processMessage(wsConn, handler)
		if err != nil {
			return err
		}
	}

	return nil
}

func processMessage(wsConn *websocket.Conn, handler interface{}) error {
	var data json.RawMessage
	var event Event

	err := wsConn.ReadJSON(&data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &event)
	if err != nil {
		return err
	}

	switch event.Type {
	case "hello":
		if r, ok := handler.(HelloReceiver); ok {
			err = r.OnHello(event)
		}
	case "message":
		if r, ok := handler.(MessageReceiver); ok {
			v := MessageEvent{}
			err = unmarshalAndInvoke(data, &v, func() error { return r.OnMessage(v) })
		}
	}

	return err
}

func unmarshalAndInvoke(data []byte, v interface{}, callback func() error) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return callback()
}
