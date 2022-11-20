package gac

import (
	"golang.org/x/oauth2"
	"net/http"
)

var (
	clientPrototype Client = nil
)

type (
	// Client represents the description for the client that does the
	// http requests towards the API
	// However this client can be substituted by anything that respects
	// this interface
	Client interface {
		SetToken(t *oauth2.Token)
		GetToken() *oauth2.Token
		IsTokenValid() bool
		DoRequest(r *http.Request) (*http.Response, error)
		RenewToken() error
	}
)

// NewWithPregeneratedClient receives a fully pregenerated and
// preset client and will use if for all operations within this
// construct.
// Please be aware that all OAuth routines must also be handled
// by that client, however the internal routines of this package
// can be used
func NewWithPregeneratedClient(cli Client) Client {
	clientPrototype = cli
	return clientPrototype
}
