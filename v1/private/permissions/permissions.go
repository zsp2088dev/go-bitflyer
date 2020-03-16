package permissions

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type Request struct{}

type Response []Permission

type Permission string

func (req *Request) Endpoint() string {
	return "/v1/me/getpermissions"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
