package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestBasicAuth(t *testing.T) {
	app := newApp()

	e := httptest.New(t, app)

	e.GET("/").Expect().Status(httptest.StatusOK)
}
