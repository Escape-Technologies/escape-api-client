package internal

import (
	"fmt"
	"net/http"
)

type TransportWithAuth struct {
	transport http.RoundTripper
	apiKey    string
}

func (t *TransportWithAuth) RoundTrip(
	req *http.Request,
) (*http.Response, error) {
	req.Header.Set("Authorization", "Key"+t.apiKey)
	tp, err := t.transport.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("error while making request: %s", err)
	}
	return tp, nil
}

func NewTransportWithAuth(
	apiKey string,
) *TransportWithAuth {
	return &TransportWithAuth{
		transport: http.DefaultTransport,
		apiKey:    apiKey,
	}
}
