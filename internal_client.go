package gac

import (
	"golang.org/x/oauth2"
	"net/http"
)

type (
	internalClient struct {
		cfg   *ClientConfig
		token *oauth2.Token
	}

	ClientConfig struct {
		Oauth2Config *oauth2.Config
		Environment  string
	}
)

func New(cfg *ClientConfig) Client {
	return NewWithPregeneratedClient(&internalClient{cfg: cfg})
}

func (ic *internalClient) SetToken(t *oauth2.Token) {
	ic.token = t
}

func (ic *internalClient) GetToken() *oauth2.Token {
	return ic.token
}

func (ic *internalClient) IsTokenValid() bool {
	token := ic.GetToken()
	return token.Valid()
}

func (ic *internalClient) DoRequest(r *http.Request) (*http.Response, error) {

	if !ic.IsTokenValid() {
		if err := ic.RenewToken(); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (ic *internalClient) RenewToken() error {
	return nil
}
