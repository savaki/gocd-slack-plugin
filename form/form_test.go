package form

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshal(t *testing.T) {
	Convey("Given a form", t, func() {
		Convey("#AsValues string fields", func() {
			form := struct {
				Name string `form:"name"`
			}{
				Name: "matt",
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Encode(), ShouldEqual, "name=matt")
			})
		})

		Convey("#AsValues int fields", func() {
			form := struct {
				Count int `form:"c"`
			}{
				Count: 123,
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Encode(), ShouldEqual, "c=123")
			})
		})

		Convey("#AsValues honors omitempty on int fields", func() {
			form := struct {
				Count int `form:"c,omitempty"`
			}{
				Count: 0,
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Encode(), ShouldEqual, "")
			})
		})

		Convey("#AsValues bool fields", func() {
			form := struct {
				Yep bool `form:"y"`
			}{
				Yep: true,
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Encode(), ShouldEqual, "y=true")
			})
		})

		Convey("#AsValues float32 fields", func() {
			form := struct {
				Floater float32 `form:"f"`
			}{
				Floater: 1.23,
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Encode(), ShouldEqual, "f=1.23")
			})
		})

		Convey("#AsValues float32 honors precision tag", func() {
			form := struct {
				Floater float32 `form:"f" precision:"3"`
			}{
				Floater: 1.234,
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Encode(), ShouldEqual, "f=1.234")
			})
		})

		Convey("#AsValues struct converts struct to json", func() {
			type Location struct {
				City  string `json:"city"`
				State string `json:"state"`
			}
			form := struct {
				Locations []Location `form:"locations"`
			}{
				Locations: []Location{
					{City: "San Francisco", State: "CA"},
				},
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Get("locations"), ShouldEqual, `[{"city":"San Francisco","state":"CA"}]`)
			})
		})

		Convey("#AsValues form.Bool", func() {
			form := struct {
				True   *bool `form:"a"`
				False  *bool `form:"b"`
				NotSet *bool `form:"c"`
			}{
				True:   True,
				False:  False,
				NotSet: nil,
			}
			params := AsValues(&form)

			Convey("Should return valid url.Values", func() {
				So(params.Get("a"), ShouldEqual, "true")
				So(params.Get("b"), ShouldEqual, "false")
				So(params.Encode(), ShouldNotContainSubstring, "c=")
			})
		})
	})
}
