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
	api *slack.Client
}

func (e *EventHandler) OnMessage(v slack.MessageEvent) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	if v.User != "" {
		go func() {
			e.api.PostMessage(slack.PostMessageReq{
				Channel:  v.Channel,
				Text:     fmt.Sprintf("I heard => %s", v.Text),
				Username: "wakka",
			})
		}()
	}

	fmt.Println(string(data))
	return nil
}

func main() {
	client, err := slack.New(os.Getenv("SLACK_TOKEN"))
	assert(err)

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
		handler := &EventHandler{api: client}
		err := client.Listen(handler)
		assert(err)
	}()
}
