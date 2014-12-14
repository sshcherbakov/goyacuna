package goyacuna

import (
	"crypto/tls"
	"net/http"
	"github.com/bndr/gopencils"
)

type Config struct {
	Url		string
	Id		string
	Secret	string
}

// Create a new API Instance and returns a Resource
// Accepts URL as parameter, and HTTP Client.
func Api(config *Config) *Instance {

	// Skip verify by default?
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	inst := Instance{ api: gopencils.Api(config.Url, client) }
	inst.SetId(config.Id)
	inst.SetSecret(config.Secret)

	return &inst

}
