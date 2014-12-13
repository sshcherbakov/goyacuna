package goyacuna

import (
	"crypto/tls"
	"net/http"
	"github.com/bndr/gopencils"
)

// Create a new API Instance and returns a Resource
// Accepts URL as parameter, and HTTP Client.
func Api(baseUrl string) *Instance {

	// Skip verify by default?
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	return &Instance{ api: gopencils.Api(baseUrl, client) }

}
