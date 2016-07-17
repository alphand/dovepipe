package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Testing Webserver Setup", t, func() {

		Convey("Webserver Port Config :3000", func() {
			So(port, ShouldEqual, ":3000")
		})

		Convey("Webserver Up and Running", func() {
			server := setupServer(port)
			So(server.Addr, ShouldEqual, ":3000")
		})

		Convey("Handle /", func() {
			req, _ := http.NewRequest("Get", "", nil)
			w := httptest.NewRecorder()
			handlerRoot().ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)
		})

	})
}
