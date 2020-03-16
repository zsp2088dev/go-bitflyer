package api

import "net/url"

type Request interface {
	Endpoint() string
	Method() string
	Query() string
	Payload() []byte
}

type API interface {
	ToURL() (*url.URL, error)
}
