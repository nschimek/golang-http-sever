package main

import (
	"testing"
)

func TestGetParamFromPathSuccess(t *testing.T) {
	actual, _ := getParamFromPath("/users/nick@test.com", "/users/")
	expected := "nick@test.com"

	if actual != expected {
		t.Errorf("actual=%q; expected=%q", actual, expected)
	}
}

func TestGetParamFromPathErrorPath(t *testing.T) {
	_, err := getParamFromPath("/users/nick@test.com", "/posts/")
	expected := "expected URL prefix not found"

	if err.Error() != expected {
		t.Errorf("actual=%q; expected=%q", err.Error(), expected)
	}
}

func TestGetParamFromPathErrorNoParam(t *testing.T) {
	_, err := getParamFromPath("/users/", "/users/")
	expected := "param value not found"

	if err.Error() != expected {
		t.Errorf("actual=%q; expected=%q", err.Error(), expected)
	}
}
