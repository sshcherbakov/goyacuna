package goyacuna

import (
	"testing"
)

func TestApiToken(t *testing.T) {

	ti := &apiTokenInput{
		secret: "79776da825eedc16abb4b0c784f7112a",
		method: "GET",
		path: 	"/api/1/wallet/get",
		query: 	"",
	}

	s := ApiToken( ti )
	t.Log(s)

	if s == "" || s[0] < '0' || s[0] > '9' {
		t.Fail()
	}

}
