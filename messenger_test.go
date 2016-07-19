package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMessenger(t *testing.T) {
	Convey("FBMessenger Unit Test", t, func() {

		Convey("Handle FB Handshake Process", func() {
			m := &Messenger{
				VerifyToken: "testtoken",
			}

			newServer := httptest.NewServer(m.Handler())
			defer newServer.Close()

			res, err := http.Get(newServer.URL + "?hub.mode=subscribe&hub.verify_token=testtoken")
			So(err, ShouldBeNil)

			content, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()

			So(string(content), ShouldEqual, m.VerifyToken)
		})

	})
}
