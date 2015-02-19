package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/savaki/gocd-slack-plugin/builtin/providers/slack"
)

func assert(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type EventHandler struct {
}

func (e *EventHandler) OnMessage(v slack.MessageEvent) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
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
		handler := &EventHandler{}
		err := client.Listen(handler)
		assert(err)

		// data, err := json.MarshalIndent(result, "", "  ")
		// assert(err)
		// fmt.Println(string(data))

		// u, err := url.Parse(result.Url)
		// assert(err)

		// target := fmt.Sprintf("%s:443", u.Host)
		// fmt.Printf("dialing %s\n", target)
		// rawConn, err := tls.Dial("tcp", target, nil)
		// assert(err)

		// wsHeaders := http.Header{
		// 	"Origin":                   {result.Url},
		// 	"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
		// }

		// wsConn, resp, err := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
		// fmt.Printf("statusCode = %d\n", resp.StatusCode)
		// v := &json.RawMessage{}

		// for {
		// 	wsConn.ReadJSON(v)
		// 	fmt.Println(string([]byte(*v)))
		// }
	}()
}
