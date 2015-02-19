package slack

import (
	"github.com/savaki/gocd-slack-plugin/form"
)

type PostMessageReq struct {
	Channel     string       `form:"channel"`
	Text        string       `form:"text"`
	Username    string       `form:"username"`
	Parse       string       `form:"parse"`
	LinkNames   int          `form:"link_names,omitempty"`
	Attachments []Attachment `form:"attachments"`
	UnfurlLinks *bool        `form:"unfurl_links"`
	UnfurlMedia *bool        `form:"unfurl_media"`
	IconUrl     string       `form:"icon_url"`
	IconEmoji   string       `form:"icon_emoji"`
}

type PostMessageResp struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Ts      string `json:"ts"`
	Channel string `json:"channel"`
}

func (c *Client) PostMessage(req PostMessageReq) (*PostMessageResp, error) {
	values := form.AsValues(req)
	resp := &PostMessageResp{}
	err := c.api("chat.postMessage", values, resp)
	return resp, err
}

type UpdateMessageReq struct {
	Ts      string `form:"ts"`
	Channel string `form:"channel"`
	Text    string `form:"text"`
}

type UpdateMessageResp struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
	Text    string `json:"ts"`
}

func (c *Client) UpdateMessage(req UpdateMessageReq) (*UpdateMessageResp, error) {
	values := form.AsValues(req)
	resp := &UpdateMessageResp{}
	err := c.api("chat.updateMessage", values, resp)
	return resp, err
}
