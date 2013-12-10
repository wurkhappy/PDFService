package main

import (
	"testing"
)

func test_ReturnString(t *testing.T){
	var params map[string]interface{}
	body := []byte(`<h1>TEST</h1>`)
	resp, _, _ := ReturnString(params, body)
	if len(resp) == 0 {
		t.Error("no string was returned")
	}
}