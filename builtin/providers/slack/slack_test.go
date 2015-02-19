package slack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApi(t *testing.T) {
	var client *Client
	var mock *MockApiFunc

	Convey("Given a Slack client", t, func() {
		mock = &MockApiFunc{}
		client = New("blah")
		client.slack = mock.Handle

		Convey("When I call #api.test", func() {
			mock.Response = `
{
    "ok": false,
    "error": "my_error",
    "args": {
        "error": "my_error"
    }
}`
			resp, err := client.ApiTest(ApiTestReq{Foo: "bar", Err: "boom!"})

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the params should be assigned correctly", func() {
				So(mock.Params, ShouldResemble, map[string]string{"foo": "bar", "error": "boom!"})
			})

			Convey("And the response should be set ok", func() {
				So(resp.Ok, ShouldBeFalse)
				So(resp.Error, ShouldEqual, "my_error")
				So(resp.Args, ShouldResemble, map[string]string{"error": "my_error"})
			})
		})

		Convey("When I call #auth.test", func() {
			mock.Response = `
{
    "ok": true,
    "url": "https:\/\/myteam.slack.com\/",
    "team": "My Team",
    "user": "cal",
    "team_id": "T12345",
    "user_id": "U12345"
}`
			resp, err := client.AuthTest()

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the response should be set ok", func() {
				So(resp.Ok, ShouldBeTrue)
				So(resp.Url, ShouldEqual, "https://myteam.slack.com/")
				So(resp.Team, ShouldEqual, "My Team")
				So(resp.User, ShouldEqual, "cal")
				So(resp.TeamId, ShouldEqual, "T12345")
				So(resp.UserId, ShouldEqual, "U12345")
			})
		})

		Convey("When I call #chat.postMessage", func() {
			mock.Response = `
{
    "ok": true,
    "ts": "1405895017.000506",
    "channel": "C024BE91L"
}`
			resp, err := client.PostMessage(PostMessageReq{})

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the response should be set ok", func() {
				So(resp.Ok, ShouldBeTrue)
				So(resp.Error, ShouldBeEmpty)
			})
		})
	})
}
