package main

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	testHeader := http.Header{
		"Host":          {"www.host.com"},
		"Content-Type":  {"application/json"},
		"Authorization": {"ApiKey fooBar"},
	}

	gotRes, gotErr := auth.GetAPIKey(testHeader)
	want := "foo"
	if gotRes != want || gotErr != nil {
		t.Errorf("got %q, wanted %q", gotRes, want)
	}
}
