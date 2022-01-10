package core

import http2 "net/http"

type http struct {
	client http2.Client
}

func (t *http) Download(url string) {
}

func newHttpClient() *http {
	client := http{}

	return &client
}
