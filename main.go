package main

import (
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
}
