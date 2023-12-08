package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 1234abc")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "1234abc"
	if key != expected {
		t.Errorf("Wrong key. Got %q, Expected %q", key, expected)
	}

}
