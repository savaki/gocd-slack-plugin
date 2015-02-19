package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	"github.com/savaki/gocd-slack-plugin/builtin/providers/slack"
)

func assert(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	client := slack.New(os.Getenv("SLACK_TOKEN"))
	func() {
		result, err := client.ApiTest(slack.ApiTestReq{Foo: "hello", Err: "world"})
		assert(err)
		fmt.Printf("%+v\n", result)
	}()

	func() {
		result, err := client.ApiTest(slack.ApiTestReq{Foo: "hello"})
		assert(err)
		fmt.Printf("%+v\n", result)
	}()

	func() {
		result, err := client.AuthTest()
		assert(err)
		fmt.Printf("%+v\n", result)
	}()

	func() {
		result, err := client.RtmStart()
		assert(err)
		fmt.Printf("%+v\n", result)

		u, err := url.Parse(result.Url)
		assert(err)

		target := fmt.Sprintf("%s:443", u.Host)
		fmt.Printf("dialing %s\n", target)
		rawConn, err := tls.Dial("tcp", target, nil)
		assert(err)

		wsHeaders := http.Header{
			"Origin":                   {result.Url},
			"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
		}

		wsConn, resp, err := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
		fmt.Printf("statusCode = %d\n", resp.StatusCode)
		v := &json.RawMessage{}

		for {
			wsConn.ReadMessage()
			wsConn.ReadJSON(v)
			fmt.Println(string([]byte(*v)))
		}
	}()
}
